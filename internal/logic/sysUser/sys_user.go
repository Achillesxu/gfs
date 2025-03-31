// Package sysUser
// Time    : 2025/3/31 16:32
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package sysUser

import (
	"context"
	"gfs/api/system/v1"
)

type sSysUser struct{}

func (s *sSysUser) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	return nil, nil
}

func (s *sSysUser) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	return nil, nil
}
