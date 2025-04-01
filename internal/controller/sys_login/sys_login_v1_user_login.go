package sys_login

import (
	"context"
	"gfs/internal/model"
	"gfs/internal/service"
	"gfs/utility"

	"github.com/gogf/gf/v2/errors/gerror"

	"gfs/api/sys_login/v1"
)

func (c *ControllerV1) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	var user *model.LoginUserRes
	if !service.Captcha().VerifyCaptcha(ctx, &model.CaptchaInput{
		Id:   req.VerifyId,
		Code: req.VerifyCode,
	}) {
		err = gerror.New("验证码输入错误")
		return
	}
	clientIp := utility.GetClientIp(ctx)
	user, err = service.SysUser().GetUserByUsernamePassword(ctx, req)
	if err != nil {
		return
	}
	err = service.SysUser().UpdateLoginInfo(ctx, user.Id, clientIp)
	if err != nil {
		return
	}

	res = &v1.UserLoginRes{
		UserInfo: user,
		Token:    "123456",
	}

	return
}
