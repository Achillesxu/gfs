// Package v1
// Time    : 2025/4/1 13:48
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package v1

import "github.com/gogf/gf/v2/frame/g"

type SetUserReq struct {
	NickName string `p:"nickName" v:"required#用户昵称不能为空"`
	Email    string `p:"email" v:"email#邮箱格式错误"` // 邮箱
	Mobile   string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	IsAdmin  int    `p:"isAdmin"` // 是否后台管理员 1:是, 0:否
	Status   uint   `p:"status"`
	Remark   string `p:"remark"`
}

type UserAddReq struct {
	g.Meta   `path:"/user/add" tags:"用户管理" method:"post" summary:"添加用户"`
	UserName string `p:"username" v:"required#用户账号不能为空"`
	Password string `p:"password" v:"required|password3#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	*SetUserReq
}

type UserAddRes struct {
}
