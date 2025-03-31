// Package system
// Time    : 2025/3/31 16:15
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package v1

import (
	"gfs/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta     `path:"/login" tags:"登录" method:"post" summary:"用户登录"`
	Username   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}

type UserLoginRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type UserLogoutReq struct {
	g.Meta `path:"/logout" tags:"登录" method:"get" summary:"退出登录"`
}

type UserLogoutRes struct {
}
