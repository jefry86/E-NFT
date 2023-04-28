package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func bannerRouter(e *gin.RouterGroup) {
	var BannerApi app.Banner
	bannerGroup := e.Group("banner")
	bannerGroup.GET("list", BannerApi.List)
}
