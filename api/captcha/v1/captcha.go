// Package v1
// Time    : 2025/3/31 17:05
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package v1

import "github.com/gogf/gf/v2/frame/g"

type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" sm:"获取验证码" tags:"验证码" dc:"获取验证码"`
}

type CaptchaRes struct {
	g.Meta `mime:"application/json"`
	Id     string `json:"id" dc:"验证码id"`
	Code   string `json:"code" dc:"验证码code(base64 string)"`
}
