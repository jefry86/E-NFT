package middleware

import (
	"github.com/gin-gonic/gin"
	"nft_platform/global"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		postForm := c.Request.PostForm
		query := c.Request.URL.Query()
		postParam := postForm.Encode()
		queryParam := query.Encode()
		url := c.Request.URL.Path
		c.Next()
		latency := time.Since(t)
		global.Logger.Info("url:" + url + " post data:" + postParam + " query data:" + queryParam + " 消耗时间:" + latency.String())
		c.Next()
	}
}
