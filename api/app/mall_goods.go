package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"nft_platform/service"
	"nft_platform/vm/request"
)

var mallGoodsService service.MallGoods

type MallGoods struct {
	UserBase
}

func (m *MallGoods) List(c *gin.Context) {
	var params request.MallGoods
	if err := c.Bind(&params); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		m.JsonParamsError(c)
		return
	}
	list, err := mallGoodsService.List(params.Sort, params.Type, params.PageNo, params.PageSize)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
		return
	}
	m.JsonSuccessWithData(c, *list)
}

// Add 发售藏品
func (m *MallGoods) Add(c *gin.Context) {

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

}
