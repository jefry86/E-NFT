package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func mallOrder(g *gin.RouterGroup) {
	var mallOrderApi app.MallOrder
	r := g.Group("/mall_order")
	{
		r.POST("/apply", mallOrderApi.Apply)
		r.GET("/info", mallOrderApi.Info)
		r.GET("/list_buy", mallOrderApi.ListForBuy)
		r.GET("/list_sale", mallOrderApi.ListForSale)
	}
}
