package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"nft_platform/service"
	"nft_platform/utils"
	"nft_platform/vm/request"
)

var mallGoodsService service.MallGoods
var platformService service.Platform

type MallGoods struct {
	UserBase
}

func (m *MallGoods) List(c *gin.Context) {
	var params request.MallGoods
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		m.JsonParamsError(c)
		return
	}
	list, err := mallGoodsService.List("", params.PlatformId, params.Sort, params.Type, params.PageNo, params.PageSize)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
		return
	}
	m.JsonSuccessWithData(c, list)
}

func (m *MallGoods) Search(c *gin.Context) {
	var params request.MallGoodsSearch
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		m.JsonParamsError(c)
		return
	}
	list, err := mallGoodsService.List(params.Keyword, params.PlatformId, params.Sort, params.Type, params.PageNo, params.PageSize)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
		return
	}
	m.JsonSuccessWithData(c, list)
}

func (m *MallGoods) Info(c *gin.Context) {
	id := utils.AtoInt(c.DefaultQuery("id", "0"))
	if id <= 0 {
		m.JsonParamsError(c)
	}

	info, err := mallGoodsService.Info(id, "")
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
	} else {
		m.JsonSuccessWithData(c, info)
	}

}

// PlatformGoodsList 藏品
func (m *MallGoods) PlatformGoodsList(c *gin.Context) {
	var params request.PlatformGoodsList
	if err := c.Bind(&params); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		m.JsonParamsError(c)
		return
	}

	if m.getUserId(c) != nil {
		return
	}

	list, err := mallPlatformServer.ApiGoodsList(m.UserId, params.PlatformId, params.PageNo, params.PageSize)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
	} else {
		m.JsonSuccessWithData(c, list)
	}

}

func (m *MallGoods) Add(c *gin.Context) {
	var params request.MallGoodsAdd
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		m.JsonParamsError(c)
		return
	}

	if m.getUserId(c) != nil {
		return
	}

	_, err := mallGoodsService.Add(params, m.UserId)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
	} else {
		m.JsonSuccess(c)
	}
}

// MyGoodsList 我的出售藏品
func (m *MallGoods) MyGoodsList(c *gin.Context) {
	if m.getUserId(c) != nil {
		return
	}
	var params request.MallGoodsList
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		m.JsonParamsError(c)
		return
	}
	list, err := mallGoodsService.ListByUserId(m.UserId, params.Status, params.PageNo, params.PageSize)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
		return
	}
	m.JsonSuccessWithData(c, list)
}
