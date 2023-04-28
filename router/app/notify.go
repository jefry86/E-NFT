package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func notify(g *gin.RouterGroup) {
	var notifyApi app.Notify
	r := g.Group("/notify")
	{
		r.POST("/pay", notifyApi.Pay) //支付通知
	}
}
