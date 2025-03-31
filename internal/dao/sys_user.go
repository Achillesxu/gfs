// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"gfs/internal/dao/internal"
)

// sysUserDao is the data access object for the table sys_user.
// You can define custom methods on it to extend its functionality as needed.
type sysUserDao struct {
	*internal.SysUserDao
}

var (
	// SysUser is a globally accessible object for table sys_user operations.
	SysUser = sysUserDao{internal.NewSysUserDao()}
)

// Add your custom methods and functionality below.
