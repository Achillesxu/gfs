// Package sysUser
// Time    : 2025/3/31 16:32
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package sysLogin

import (
	"context"
	"gfs/api/sys_login/v1"
	"gfs/internal/service"
)

type sSysLogin struct{}

func init() {
	service.RegisterSysLogin(New())
}

func New() *sSysLogin {
	return &sSysLogin{}
}

func (s *sSysLogin) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {

	return nil, nil
}

func (s *sSysLogin) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {

	return nil, nil
}
