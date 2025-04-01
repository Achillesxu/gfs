package captcha

import (
	"context"

	"gfs/api/captcha/v1"
	"gfs/internal/service"
)

func (c *ControllerV1) Captcha(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	s := service.Captcha()
	res, err = s.Captcha(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
