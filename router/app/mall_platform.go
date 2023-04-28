package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func mallPlatform(g *gin.RouterGroup) {
	var mallPlatformApi app.MallPlatform
	r := g.Group("/mall_platform")
	{
		r.GET("/list", mallPlatformApi.List)
	}
}
