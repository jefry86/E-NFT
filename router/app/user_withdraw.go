package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func userWithdraw(g *gin.RouterGroup) {
	var userWithdrawApi app.UserWithdraw
	r := g.Group("/user_withdraw")
	{
		r.GET("/list", userWithdrawApi.List) //提现记录
		r.POST("/add", userWithdrawApi.Add)  //提现
	}
}
