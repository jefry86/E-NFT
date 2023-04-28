package app

import "github.com/gin-gonic/gin"

func AppRouter(e *gin.Engine) {
	routerGroup := e.Group("app")
	bannerRouter(routerGroup)
	mallGoods(routerGroup)
	mallOrder(routerGroup)
	mallPlatform(routerGroup)
	mallUserGoods(routerGroup)
	notify(routerGroup)
	userBalance(routerGroup)
	userBank(routerGroup)
	userWithdraw(routerGroup)
	users(routerGroup)
}
