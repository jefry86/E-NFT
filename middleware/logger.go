package middleware

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		err := c.Request.ParseForm()
		if err != nil {
			global.SLogger.Errorf("err:%s", err.Error())
			return
		}

		query := c.Request.URL.Query()
		postForm := c.Request.PostForm
		postParam := postForm.Encode()
		queryParam := query.Encode()
		url := c.Request.URL.Path
		c.Next()
		latency := time.Since(t)
		global.Logger.Info("url:" + url + " post data:" + postParam + ", query data:" + queryParam + " 消耗时间:" + latency.String())
		c.Next()
	}
}
