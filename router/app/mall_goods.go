package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func mallGoods(g *gin.RouterGroup) {
	var mallGoodsApi app.MallGoods
	r := g.Group("/mall_goods")
	{
		r.GET("/list", mallGoodsApi.List)
		r.GET("/search", mallGoodsApi.Search)
		r.GET("/info", mallGoodsApi.Info)
		r.POST("/add", mallGoodsApi.Add)
		r.GET("/platform_goods_list", mallGoodsApi.PlatformGoodsList)
		r.GET("/my_goods_list", mallGoodsApi.MyGoodsList)
	}
}
