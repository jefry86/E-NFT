package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func public(g *gin.RouterGroup) {
	var publicApi app.Public
	r := g.Group("/public")
	{
		r.GET("/tcloud_policy", publicApi.TCloudPolicy)
		r.GET("/tcloud_sts", publicApi.TCloudSTS)
	}
}
