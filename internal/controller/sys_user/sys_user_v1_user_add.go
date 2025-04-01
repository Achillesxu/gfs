package sys_user

import (
	"context"
	"gfs/internal/service"

	"gfs/api/sys_user/v1"
)

func (c *ControllerV1) UserAdd(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error) {
	err = service.SysUser().AddUser(ctx, req)
	return &v1.UserAddRes{}, err
}
