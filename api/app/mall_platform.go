package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api"
	"nft_platform/global"
	"nft_platform/service"
	"nft_platform/vm/request"
)

var mallPlatformServer service.MallPlatform

type MallPlatform struct {
	api.Api
}

func (m *MallPlatform) List(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBind(&page); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		m.JsonParamsError(c)
		return
	}

	list, err := mallPlatformServer.List(page.PageNo, page.PageSize)
	if err != nil {
		m.JsonErrorWithMsg(c, err.Error())
	} else {
		m.JsonSuccessWithData(c, list)
	}
}
