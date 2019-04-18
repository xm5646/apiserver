/**
 * 功能描述: 工具类
 * @Date: 2019-04-17
 * @author: lixiaoming
 */
package util

import "github.com/gin-gonic/gin"

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
