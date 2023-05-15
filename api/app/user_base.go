package app

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api"
	"nft_platform/service"
)

type UserBase struct {
	api.Api
	UserId string
}

func (u *UserBase) getUserId(c *gin.Context) error {
	var userService service.Users
	var err error
	u.UserId, err = userService.GetUserByToken(c.GetHeader("Authorization"))
	if err != nil {
		u.JsonErrorWithMsg(c, err.Error())
		return err
	}
	return nil
}
