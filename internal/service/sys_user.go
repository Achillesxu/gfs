// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	loginV1 "gfs/api/sys_login/v1"
	userV1 "gfs/api/sys_user/v1"
	"gfs/internal/model"

	"github.com/gogf/gf/v2/container/gset"
)

type (
	ISysUser interface {
		NotCheckAuthAdminIds(ctx context.Context) *gset.Set
		AddUser(ctx context.Context, req *userV1.UserAddReq) (err error)
		UserNameOrMobileExists(ctx context.Context, userName string, mobile string, id ...int64) (err error)
		GetUserByUsernamePassword(ctx context.Context, req *loginV1.UserLoginReq) (user *model.LoginUserRes, err error)
		// GetUserByUsername 通过用户名获取用户信息
		GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error)
		UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
