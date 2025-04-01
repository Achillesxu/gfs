// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_user

import (
	"context"

	"gfs/api/sys_user/v1"
)

type ISysUserV1 interface {
	UserAdd(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error)
}
