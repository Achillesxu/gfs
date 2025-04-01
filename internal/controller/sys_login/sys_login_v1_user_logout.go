package sys_login

import (
	"context"

	"gfs/api/sys_login/v1"
)

func (c *ControllerV1) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	res = &v1.UserLogoutRes{}
	return
}
