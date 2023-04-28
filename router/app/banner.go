package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func bannerRouter(g *gin.RouterGroup) {
	var BannerApi app.Banner
	bannerGroup := g.Group("banner")
	bannerGroup.GET("list", BannerApi.List)
}
