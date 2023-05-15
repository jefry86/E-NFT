package middleware

import (
	"github.com/gin-gonic/gin"
	"nft_platform/api"
	"nft_platform/global"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var R api.Api
				global.SLogger.Error("系统出现异常，err:%s", err)
				R.JsonError(ctx)
			}
		}()
	}
}
