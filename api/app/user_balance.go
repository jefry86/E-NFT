package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/service"
)

var userBalanceService service.UserBalance

type UserBalance struct {
	UserBase
}

// List 余额流水
func (b *UserBalance) List(c *gin.Context) {
	if b.getUserId(c) != nil {
		return
	}

	list, err := userBankService.List(b.UserId)
	if err != nil {
		b.JsonErrorWithMsg(c, err.Error())
	} else {
		b.JsonSuccessWithData(c, list)
	}
}
