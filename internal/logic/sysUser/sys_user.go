// Package sysUser
// Time    : 2025/4/1 14:24
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package sysUser

import (
	"context"
	"fmt"
	loginV1 "gfs/api/sys_login/v1"
	userV1 "gfs/api/sys_user/v1"
	"gfs/internal/dao"
	"gfs/internal/model"
	"gfs/internal/model/do"
	"gfs/internal/model/entity"
	"gfs/internal/service"
	"gfs/utility"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type sSysUser struct{}

func init() {
	service.RegisterSysUser(New())
}

func New() *sSysUser {
	return &sSysUser{}
}

func (s *sSysUser) NotCheckAuthAdminIds(ctx context.Context) *gset.Set {
	ids := g.Cfg().MustGet(ctx, "system.notCheckAuthAdminIds")
	if !g.IsNil(ids) {
		return gset.NewFrom(ids)
	}
	return gset.New()
}

func (s *sSysUser) AddUser(ctx context.Context, req *userV1.UserAddReq) (err error) {
	if err = s.UserNameOrMobileExists(ctx, req.UserName, req.Mobile); err != nil {
		return
	}
	userSalt := grand.S(10)
	req.Password = utility.EncryptPassword(req.Password, userSalt)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, e := dao.SysUser.Ctx(ctx).TX(tx).InsertAndGetId(do.SysUser{
			UserName:     req.UserName,
			Mobile:       req.Mobile,
			UserNickname: req.NickName,
			UserPassword: req.Password,
			UserSalt:     userSalt,
			Status:       req.Status,
			Email:        req.Email,
			IsAdmin:      req.IsAdmin,
			Remark:       req.Remark,
		})
		if e != nil {
			e = gerror.Wrap(e, "添加用户失败")
			g.Log().Error(ctx, e.Error())
		}
		return e
	})
	return
}

func (s *sSysUser) UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) (err error) {
	user := (*entity.SysUser)(nil)
	m := dao.SysUser.Ctx(ctx)
	if len(id) > 0 {
		m = m.Where(dao.SysUser.Columns().Id+" != ", id)
	}
	m = m.Where(fmt.Sprintf("%s='%s' OR %s='%s'",
		dao.SysUser.Columns().UserName,
		userName,
		dao.SysUser.Columns().Mobile,
		mobile))
	err = m.Limit(1).Scan(&user)

	if err != nil {
		err = gerror.Wrap(err, "获取用户信息失败")
		g.Log().Error(ctx, err)
		return
	}

	if user == nil {
		err = nil
		return
	}
	if user.UserName == userName {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "用户名已存在")
		return
	}
	if user.Mobile == mobile {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "手机号已存在")
		return
	}
	return
}

func (s *sSysUser) GetUserByUsernamePassword(ctx context.Context, req *loginV1.UserLoginReq) (user *model.LoginUserRes, err error) {
	user, err = s.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return
	}
	if g.IsNil(user) {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "账号不存在")
		return
	}
	// 验证密码
	if utility.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "账号密码错误")
		return
	}
	// 账号状态
	if user.Status == 0 {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "账号已被冻结")
		return
	}
	return
}

// GetUserByUsername 通过用户名获取用户信息
func (s *sSysUser) GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error) {
	user = &model.LoginUserRes{}
	err = dao.SysUser.Ctx(ctx).Fields(user).Where(dao.SysUser.Columns().UserName, userName).Scan(user)
	if err != nil {
		err = gerror.Wrap(err, "服务内部错误")
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (s *sSysUser) UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error) {
	_, e := dao.SysUser.Ctx(ctx).WherePri(id).Unscoped().Update(g.Map{
		dao.SysUser.Columns().LastLoginIp:   ip,
		dao.SysUser.Columns().LastLoginTime: gtime.Now(),
	})
	if e != nil {
		err = gerror.Wrap(e, "更新用户登录信息失败")
		return
	}
	return
}
