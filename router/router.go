/**
 * 功能描述: 路由管理
 * @Date: 2019-04-14
 * @author: lixiaoming
 */
package router

import (
	_ "apiserver/docs"

	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// 加载中间件, 路由,返回gin引擎
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API server is running.")
	})
	// 加载swagger api 文档
	// 重新生成文档  swag init
	// 访问地址/swagger/index.html
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	loginGroup := g.Group("/login")
	loginGroup.POST("/phone", user.PhoneLogin)
	loginGroup.POST("/phone/check/:phoneNumber", user.CheckPhoneIsRegistered)

	userGroup := g.Group("/v1/user")
	userGroup.Use(middleware.AuthMiddleware())
	{
	}

	return g
}
