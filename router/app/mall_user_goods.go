package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func mallUserGoods(g *gin.RouterGroup) {
	var mallUserGoodsApi app.MallUserGoods
	r := g.Group("/mall_user_goods")
	{
		r.GET("/list", mallUserGoodsApi.GoodsList) //藏品列表
		r.GET("/info", mallUserGoodsApi.Info)      //藏品详情
	}
}
