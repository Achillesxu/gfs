// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gfs/api/sys_login/v1"
)

type (
	ISysLogin interface {
		UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
		UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error)
	}
)

var (
	localSysLogin ISysLogin
)

func SysLogin() ISysLogin {
	if localSysLogin == nil {
		panic("implement not found for interface ISysLogin, forgot register?")
	}
	return localSysLogin
}

func RegisterSysLogin(i ISysLogin) {
	localSysLogin = i
}
