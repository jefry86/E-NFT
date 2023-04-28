package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"nft_platform/service"
	"nft_platform/utils"
	"nft_platform/vm/request"
)

var mallUserGoodsService service.MallUserGoods

type MallUserGoods struct {
	UserBase
}

func (u *MallUserGoods) GoodsList(c *gin.Context) {

	var params request.MallUserGoods
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		u.JsonParamsError(c)
		return
	}

	if u.getUserId(c) != nil {
		return
	}

	list, err := mallUserGoodsService.GoodsList(u.UserId, params.Type, params.PageNo, params.PageSize)
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccessWithData(c, *list)
	}
}

func (u *MallUserGoods) Info(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		u.JsonParamsError(c)
		return
	}
	if u.getUserId(c) != nil {
		return
	}

	info, err := mallUserGoodsService.GoodsInfo(u.UserId, utils.AtoInt(id))
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccessWithData(c, *info)
	}

}
