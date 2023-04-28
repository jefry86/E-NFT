package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"nft_platform/service"
	"nft_platform/vm/request"
)

var userWithdrawService service.UserWithdraw

type UserWithdraw struct {
	UserBase
}

func (u *UserWithdraw) List(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBind(&page); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		u.JsonParamsError(c)
		return
	}

	if u.getUserId(c) != nil {
		return
	}

	list, err := userWithdrawService.List(u.UserId, page.PageNo, page.PageSize)
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccessWithData(c, list)
	}

}

func (u *UserWithdraw) Add(c *gin.Context) {
	var params request.UserWithdrawAdd
	if err := c.ShouldBind(&params); err != nil {
		global.SLogger.Errorf("param error:%s", err)
		u.JsonParamsError(c)
		return
	}

	if u.getUserId(c) != nil {
		return
	}

	_, err := userWithdrawService.Add(u.UserId, params.BankId, params.Amount)
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
	} else {
		u.JsonSuccess(c)
	}
}
