/**
 * 功能描述: Middleware 为request header添加requestID
 * @Date: 2019-04-17
 * @author: lixiaoming
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 查看请求是否包含requestID, 如果有就直接使用
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		// 在context 中设置requestID
		c.Set("X-Request-Id", requestId)

		// 在header中设置requestID
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()

	}
}
