package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"nft_platform/vm/request"
)

type Notify struct {
}

func (n *Notify) Pay(c *gin.Context) {
	var params request.NotifyPay
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		c.String(200, "%s", "fail")
		return
	}

	var err error
	var payChannelId uint
	switch params.PayChannel {
	case "wxpay":
		payChannelId = 1
	case "alipay":
		payChannelId = 2
	case "applepay":
		payChannelId = 3
	default:
		err = fmt.Errorf("支付渠道有误！")
	}
	if err != nil {
		global.SLogger.Errorf("pay channel error:%s", err)
		c.String(200, "%s", "fail")
	}

	_, err = mallOrderService.Notify(params.OrderId, params.PayTime, payChannelId, params.Amount)
	if err != nil {
		c.String(200, "%s", "fail")
	} else {
		c.String(200, "%s", "success")
	}
}
