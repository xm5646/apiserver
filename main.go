package main

import (
	"apiserver/model"
	"apiserver/pkg/config"
	"apiserver/router"
	"apiserver/router/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

// @title  拼车车App后台服务器API文档
// @version 0.2
// @description 拼车车App后台服务器,主要提供容器集群调度服务.
// @description <a href='?docExpansion=none'>折叠</a>  <a href='index.html'>展开</a>
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// 解析命令行参数
	pflag.Parse()

	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	// 创建Gin引擎
	gin.SetMode(viper.GetString("server.runmode"))
	g := gin.New()
	var middlewares []gin.HandlerFunc
	middlewares = append(middlewares, middleware.RequestId())
	middlewares = append(middlewares, middleware.Logging())

	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// 配置TLS
	cert := viper.GetString("server.tls.cert")
	key := viper.GetString("server.tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("server.tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("server.tls.addr"), cert, key, g).Error())
		}()
	}
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("server.addr"))
	log.Infof(http.ListenAndServe(viper.GetString("server.addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("server.max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("server.url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// sleep
		log.Info("Waiting for the router, retry in i second")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router. ")
}
