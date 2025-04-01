// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta        `orm:"table:sys_user, do:true"`
	Id            interface{} //
	UserName      interface{} // 用户名
	UserNickname  interface{} // 用户昵称
	UserPassword  interface{} // 登录密码;cmf_password加密
	UserSalt      interface{} // 加密盐
	Status        interface{} // 用户状态;0:禁用,1:正常,2:未验证
	Mobile        interface{} // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	Email         interface{} // 用户登录邮箱
	IsAdmin       interface{} // 是否后台管理员 1:是, 0:否
	Remark        interface{} // 备注
	LastLoginIp   interface{} // 最后登录ip
	LastLoginTime *gtime.Time // 最后登录时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
