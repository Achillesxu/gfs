// Package captcha
// Time    : 2025/3/31 17:27
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package captcha

import (
	"context"
	v1 "gfs/api/captcha/v1"
	"gfs/internal/model"
	"gfs/internal/service"
	"gfs/utility"
	"github.com/mojocn/base64Captcha"
)

type sCaptcha struct {
	Driver *base64Captcha.DriverString
	Store  base64Captcha.Store
}

func init() {
	service.RegisterCaptcha(New())
}

func New() *sCaptcha {
	return &sCaptcha{
		Driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      50,
			ShowLineOptions: 20,
			Length:          6,
			Source:          "abcdefghjkmnpqrstuvwxyz23456789ABCDEFGHJKMNPQRSTUVWXYZ",
			Fonts:           []string{"chromohv.ttf"},
		},
		Store: utility.NewCaptchaRedisStore(),
	}
}

func (s *sCaptcha) Captcha(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	driver := s.Driver.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, s.Store)
	cId, cB64s, _, err := c.Generate()

	if err != nil {
		return nil, err
	}

	res = &v1.CaptchaRes{
		Id:   cId,
		Code: cB64s,
	}

	return
}

func (s *sCaptcha) VerifyCaptcha(ctx context.Context, in *model.CaptchaInput) bool {
	if in.Id != "" && in.Code != "" {
		v := s.Store.Verify(in.Id, in.Code, true)
		return v
	} else {
		return false
	}
}
