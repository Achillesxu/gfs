// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gfs/api/captcha/v1"
	"gfs/internal/model"
)

type (
	ICaptcha interface {
		Captcha(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error)
		VerifyCaptcha(ctx context.Context, in *model.CaptchaInput) bool
	}
)

var (
	localCaptcha ICaptcha
)

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
