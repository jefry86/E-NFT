package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api/app"
)

func userBank(g *gin.RouterGroup) {
	var userBankApi app.UserBank
	r := g.GET("/user_bank")
	{
		r.GET("/list", userBankApi.List)        //银行卡列表
		r.POST("/add", userBankApi.Add)         //添加银行卡
		r.DELETE("/delete", userBankApi.Delete) //移除银行卡
	}
}
