package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"nft_platform/service"
	"nft_platform/utils"
	"nft_platform/vm/response"
)

var userBankService service.UserBank

type UserBank struct {
	UserBase
}

// List 我的银行卡列表
func (u *UserBank) List(c *gin.Context) {
	if u.getUserId(c) != nil {
		return
	}
	list, err := userBankService.List(u.UserId)
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccessWithData(c, list)
	}
}

// Add 添加银行卡
func (u *UserBank) Add(c *gin.Context) {
	var params response.UserBankInfo
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Warnf("请求参数有误,err:%s", err)
		u.JsonParamsError(c)
		return
	}

	if u.getUserId(c) != nil {
		return
	}

	_, err := userBankService.Add(params, u.UserId)
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccess(c)
	}

}

// Del 移除银行卡
func (u *UserBank) Del(c *gin.Context) {
	if u.getUserId(c) != nil {
		return
	}
	id := utils.AtoInt(c.Param("id"))
	if id == 0 {
		u.JsonParamsError(c)
		return
	}
	_, err := userBankService.Delete(id, u.UserId)
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccess(c)
	}
}
