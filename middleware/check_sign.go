package middleware

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"nft_platform/api"
	"nft_platform/global"
	"nft_platform/utils"
	"strings"
)

func CheckSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		//只有 /app 的接口才有sign
		appPath := "/app"
		uri := c.Request.RequestURI
		if !strings.Contains(uri, appPath) {
			return
		}

		timestamp := c.GetHeader("timestamp")
		sign := c.GetHeader("sign")
		appKey := c.Query("app_key")
		var R api.Api
		if timestamp == "" || sign == "" || appKey == "" {
			global.SLogger.Warnf("timestamp:%s sign:%s app_key:%s 为空", timestamp, sign, appKey)
			R.JsonWithCode(c, global.ParamErr)
			c.Abort()
			return
		} else if appKey != global.Conf.Server.AppKey {
			global.SLogger.Warnf("app key:%s 错误", appKey)
			R.JsonWithCode(c, global.ParamErr)
			c.Abort()
			return
		}

		queryParam := c.Request.URL.Query()
		postParam := c.Request.PostForm
		var params url.Values
		if len(queryParam) > 0 {
			params = queryParam
		} else if len(postParam) > 0 {
			params = postParam
		}
		params.Add("timestamp", timestamp)

		data := make(map[string]interface{}, 0)
		for key, val := range params {
			data[key] = val[0]
		}
		toSign := utils.ToSign(data, global.Conf.Server.AppSecret)
		if sign != toSign {
			global.SLogger.Warnf("验证签名失败,sign:%s toSign:%s", sign, toSign)
			R.JsonWithCode(c, global.SignErr)
			c.Abort()
			return
		}
		c.Next()
	}
}
