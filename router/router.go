package router

import (
	"github.com/gin-gonic/gin"
	"nft_platform/router/admin"
	"nft_platform/router/app"
)

func NewRouter(e *gin.Engine) {
	noRoute(e)
	app.AppRouter(e)
	admin.AdminRouter(e)
}

func noRoute(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.JSON(404, "NOT FOUND")
	})
}
