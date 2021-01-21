/**
 * 功能描述: API身份认证
 * @Date: 2019-04-17
 * @author: lixiaoming
 */
package middleware

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xm5646/log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		contextToken, err := token.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("user_id", contextToken.ID)
		c.Set("user_name", contextToken.Username)
		log.Infof("verify the token success.")
		c.Next()
	}
}

// 验证websocket是否带有param token
func WebSocketAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Infof("verify the websocket token.")
		tokenStr := c.Request.Header.Get("Sec-WebSocket-Protocol")
		// 从配置文件加载jwt secret
		secret := viper.GetString("server.jwt_secret")
		contextToken, err := token.Parse(tokenStr, secret)
		if err != nil {
			log.Errorf(err, "failed to parse the websocket token in auth middleware.")
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("user_id", contextToken.ID)
		c.Set("user_name", contextToken.Username)
		log.Infof("verify the websocket token success.")
		c.Next()
	}
}
