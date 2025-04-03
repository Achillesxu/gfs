package cmd

import (
	"context"
	"gfs/internal/consts"
	"gfs/internal/controller/captcha"
	"gfs/internal/controller/sys_login"
	"gfs/internal/controller/sys_user"
	"gfs/internal/service"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
			g.Log().Info(ctx, "GFS\t", "Version:", consts.Version)
			s := g.Server()
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					captcha.NewV1(),
					sys_login.NewV1(),
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					_ = service.GfToken().Middleware(group)
					group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
					group.Bind(
						sys_user.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
