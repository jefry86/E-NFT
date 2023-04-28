package app

import "github.com/gin-gonic/gin"

func AppRouter(e *gin.Engine) {
	routerGroup := e.Group("app")
	bannerRouter(routerGroup)
}
