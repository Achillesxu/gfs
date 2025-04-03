package sys_login

import (
	"context"
	"gfs/internal/service"
	"github.com/gogf/gf/v2/frame/g"

	"gfs/api/sys_login/v1"
)

func (c *ControllerV1) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	err = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}
