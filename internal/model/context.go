// Package model
// Time    : 2025/4/3 10:36
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package model

type Context struct {
	User *ContextUser // User in context.
}

type ContextUser struct {
	*LoginUserRes
}
