// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id            uint64      `json:"id"            orm:"id"              description:""`                             //
	UserName      string      `json:"userName"      orm:"user_name"       description:"用户名"`                          // 用户名
	UserNickname  string      `json:"userNickname"  orm:"user_nickname"   description:"用户昵称"`                         // 用户昵称
	UserPassword  string      `json:"userPassword"  orm:"user_password"   description:"登录密码;cmf_password加密"`          // 登录密码;cmf_password加密
	UserSalt      string      `json:"userSalt"      orm:"user_salt"       description:"加密盐"`                          // 加密盐
	UserStatus    uint        `json:"userStatus"    orm:"user_status"     description:"用户状态;0:禁用,1:正常,2:未验证"`         // 用户状态;0:禁用,1:正常,2:未验证
	Mobile        string      `json:"mobile"        orm:"mobile"          description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserEmail     string      `json:"userEmail"     orm:"user_email"      description:"用户登录邮箱"`                       // 用户登录邮箱
	Remark        string      `json:"remark"        orm:"remark"          description:"备注"`                           // 备注
	IsAdmin       int         `json:"isAdmin"       orm:"is_admin"        description:"是否后台管理员 1 是  0   否"`           // 是否后台管理员 1 是  0   否
	LastLoginIp   string      `json:"lastLoginIp"   orm:"last_login_ip"   description:"最后登录ip"`                       // 最后登录ip
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"last_login_time" description:"最后登录时间"`                       // 最后登录时间
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`                         // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`                         // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:"删除时间"`                         // 删除时间
}
