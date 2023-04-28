package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"nft_platform/service"
	"nft_platform/vm/request"
)

var mallOrderService service.MallOrder

type MallOrder struct {
	UserBase
}

// Apply 下单
func (m *MallOrder) Apply(c *gin.Context) {
	var params request.MallOrderApply
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		m.JsonParamsError(c)
		return
	}

	if m.getUserId(c) == nil {
		return
	}

	orderId, err := mallOrderService.Apply(m.UserId, params.GoodsId, params.Num)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
	} else {
		m.JsonSuccessWithData(c, map[string]string{
			"order_id": orderId,
		})
	}
}

// Info 订单详细
func (m *MallOrder) Info(c *gin.Context) {
	orderId, ok := c.GetQuery("order_id")
	if !ok {
		global.SLogger.Warn("请求参数有误,ok:false")
		m.JsonParamsError(c)
		return
	} else if orderId == "" {
		global.SLogger.Warn("请求参数有误,order id 为空")
		m.JsonParamsError(c)
		return
	}
	if m.getUserId(c) != nil {
		return
	}

	info, err := mallOrderService.Info(m.UserId, orderId)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
		return
	}
	m.JsonSuccessWithData(c, info)

}

// ListForBuy 购买订单列表
func (m *MallOrder) ListForBuy(c *gin.Context) {
	if m.getUserId(c) != nil {
		return
	}
	var params request.MallOrderList
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		m.JsonParamsError(c)
		return
	}

	if err := m.getUserId(c); err != nil {
		return
	}

	list, err := mallOrderService.ListForBuy(m.UserId, params.Status, params.PageNo, params.PageSize)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
	} else {
		m.JsonSuccessWithData(c, list)
	}
}

// ListForSale 出售订单列表
func (m *MallOrder) ListForSale(c *gin.Context) {
	if m.getUserId(c) != nil {
		return
	}
	var params request.MallOrderList
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		m.JsonParamsError(c)
		return
	}

	if err := m.getUserId(c); err != nil {
		return
	}

	list, err := mallOrderService.ListForSale(m.UserId, params.Status, params.PageNo, params.PageSize)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
	} else {
		m.JsonSuccessWithData(c, list)
	}
}
