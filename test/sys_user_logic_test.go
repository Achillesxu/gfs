// Package test
// Time    : 2025/4/1 13:43
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package test

import (
	v1 "gfs/api/sys_user/v1"
	"gfs/internal/logic/sysUser"
	"gfs/internal/service"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestAddUser(t *testing.T) {
	service.RegisterSysUser(sysUser.New())
	err := service.SysUser().AddUser(gctx.New(), &v1.UserAddReq{
		UserName: "admin",
		Password: "Cyber@2077",
		SetUserReq: &v1.SetUserReq{
			NickName: "admin",
			Email:    "achilles.xu@outlook.com",
			Mobile:   "18682193124",
			IsAdmin:  1,
			Status:   1,
			Remark:   "",
		},
	})

	if err != nil {
		panic(err)
	}

}
