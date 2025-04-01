// Package utility
// Time    : 2025/3/31 17:19
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package utility

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mojocn/base64Captcha"
)

type CaptchaRedisStore struct {
	Expiration int64
	PreKey     string
	Context    context.Context
}

func NewCaptchaRedisStore() *CaptchaRedisStore {
	return &CaptchaRedisStore{
		Expiration: 180,
		PreKey:     "CAPTCHA:",
		Context:    context.Background(),
	}
}

func (rs *CaptchaRedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *CaptchaRedisStore) Set(id string, value string) error {
	var redis = g.Redis()
	err := redis.SetEX(rs.Context, rs.PreKey+id, value, rs.Expiration)

	if err != nil {
		return err
	}
	return nil
}

func (rs *CaptchaRedisStore) Get(key string, clear bool) string {
	var redis = g.Redis()
	val, err := redis.Get(rs.Context, rs.PreKey+key)
	if err != nil {
		g.Log().Errorf(gctx.New(), "RedisCaptchaStore Get(%s, %t) %#v", key, clear, err)
		return ""
	}

	if clear {
		_, err := redis.Del(rs.Context, key)
		if err != nil {
			g.Log().Errorf(gctx.New(), "RedisCaptchaStore Get(%s, %t) Del %#v", key, clear, err)
			return ""
		}
	}
	return val.String()
}

func (rs *CaptchaRedisStore) Verify(id, answer string, clear bool) bool {
	v := rs.Get(id, clear)
	return v == answer
}
