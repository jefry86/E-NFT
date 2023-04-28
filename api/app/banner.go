package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api"
	"nft_platform/model"
	"nft_platform/utils"
)

type Banner struct {
	api.Api
}

func (b *Banner) List(c *gin.Context) {
	t := c.DefaultQuery("type", "1")
	bannerModel := model.NftBanner{}
	list, err := bannerModel.List(utils.AtoInt(t))
	if err != nil {
		b.JsonError(c)
	} else {
		b.JsonSuccessWithData(c, *list)
	}
}
