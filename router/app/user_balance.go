package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func userBalance(g *gin.RouterGroup) {
	var userBalanceApi app.UserBalance
	r := g.Group("/user_balance")
	{
		r.GET("/list", userBalanceApi.List) //余额流水
	}
}
