package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func users(g *gin.RouterGroup) {
	var userApi app.Users
	r := g.Group("/user")
	{
		r.POST("/login_by_mobile", userApi.LoginByMobile) //通过手机号码登录
		r.POST("/auth", userApi.Auth)                     //实名认证
		r.DELETE("/logout", userApi.Logout)               //退出登录
		r.POST("/change_mobile", userApi.ChangMobile)     //修改手机号码
		r.POST("/send_sms_code", userApi.SendSMS)         //发送验证码
		r.GET("/info", userApi.Info)                      //用户信息
	}
}
