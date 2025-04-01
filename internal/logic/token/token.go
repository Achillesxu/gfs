// Package token
// Time    : 2025/4/1 19:35
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package token

import (
	"gfs/internal/model"
	"gfs/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-token/gftoken"
)

type sToken struct {
	*gftoken.GfToken
}

func New() *sToken {
	var (
		ctx = gctx.New()
		opt *model.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
		fun gftoken.OptionFunc
	)
	if err != nil {
		g.Log().Panic(ctx, err)
	}

	if opt.CacheModel == "redis" {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	return &sToken{
		GfToken: gftoken.NewGfToken(
			gftoken.WithCacheKey(opt.CacheKey),
			gftoken.WithTimeout(opt.Timeout),
			gftoken.WithMaxRefresh(opt.MaxRefresh),
			gftoken.WithMultiLogin(opt.MultiLogin),
			gftoken.WithExcludePaths(opt.ExcludePaths),
			fun,
		),
	}
}

func init() {
	service.RegisterGToken(New())
}
