package sys_login

import (
	"context"
	"gfs/internal/model"
	"gfs/internal/service"
	"gfs/utility"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"

	"gfs/api/sys_login/v1"
)

func (c *ControllerV1) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	var (
		user  *model.LoginUserRes
		token string
	)
	if !service.Captcha().VerifyCaptcha(ctx, &model.CaptchaInput{
		Id:   req.VerifyId,
		Code: req.VerifyCode,
	}) {
		err = gerror.New("验证码输入错误")
		return
	}
	clientIp := utility.GetClientIp(ctx)
	userAgent := utility.GetUserAgent(ctx)
	user, err = service.SysUser().GetUserByUsernamePassword(ctx, req)
	if err != nil {
		return
	}
	err = service.SysUser().UpdateLoginInfo(ctx, user.Id, clientIp)
	if err != nil {
		return
	}

	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) +
			gmd5.MustEncryptString(user.UserPassword+clientIp+userAgent)
	}
	token, err = service.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	user.UserPassword = ""
	res = &v1.UserLoginRes{
		UserInfo: user,
		Token:    token,
	}

	return
}
