package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"nft_platform/service"
	"nft_platform/utils"
	"nft_platform/vm/request"
)

var smsService service.Sms
var userService service.Users

type Users struct {
	UserBase
}

// LoginByMobile 登录
func (u *Users) LoginByMobile(c *gin.Context) {
	var params request.LoginByMobile
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		u.JsonParamsError(c)
		return
	}
	result, err := userService.LoginByMobile(params.Mobile, params.Code)
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
		return
	}
	u.JsonSuccessWithData(c, result)
}

// Auth 实名认证
func (u *Users) Auth(c *gin.Context) {
	var params request.Auth
	if err := c.Bind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		u.JsonParamsError(c)
		return
	}

	userId, err := userService.GetUserByToken(c.GetHeader("token"))
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
		return
	}

	if err := userService.Auth(userId, params.RealName, params.CardNo); err != nil {
		u.JsonErrorWithMsg(c, err.Error())
		return
	}
	u.JsonSuccess(c)
}

// Logout 退出
func (u *Users) Logout(c *gin.Context) {
	if err := userService.Logout(c.GetHeader("token")); err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccess(c)
	}
}

// ChangMobile 修改手机号码
func (u *Users) ChangMobile(c *gin.Context) {
	var params request.ChangeMobile
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		u.JsonParamsError(c)
		return
	}
	if err := u.getUserId(c); err != nil {
		return
	}
	var err error
	var ok bool
	switch params.Step {
	case 1: //验证原手机
		ok, err = userService.VerifyOldMobile(u.UserId, params.Code)
	case 2: //验证新手机
		ok, err = userService.ChangeMobile(u.UserId, params.Mobile, params.Code)
	default:
		u.JsonParamsError(c)
	}

	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else if ok {
		u.JsonSuccess(c)
	}
}

// SendSMS 发送验证码
func (u *Users) SendSMS(c *gin.Context) {
	var params request.SendSMS
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		u.JsonParamsError(c)
		return
	}

	if params.T == "2" && u.getUserId(c) != nil {
		return
	}
	if err := smsService.SendCode(u.UserId, params.Mobile, c.ClientIP(), utils.AtoInt(params.T)); err != nil {
		u.JsonErrorWithMsg(c, err.Error())
		return
	}
	u.JsonSuccess(c)
}

// Info 用户信息
func (u *Users) Info(c *gin.Context) {
	if u.getUserId(c) != nil {
		return
	}

	info, err := userService.Info(u.UserId)
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccessWithData(c, info)
	}
}
