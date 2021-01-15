/**
 * 功能描述: 日志中间件,记录请求
 * @Date: 2019-04-17
 * @author: lixiaoming
 */
package middleware

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/willf/pad"
	"github.com/xm5646/log"
	"io/ioutil"
	"regexp"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 日志中间件, 记录每一次请求
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		path := c.Request.URL.Path

		// 只对正则匹配的请求进行记录
		reg := regexp.MustCompile("(/v1/user|/login)")
		if !reg.MatchString(path) {
			return
		}

		// 跳过健康状态检查的请求
		if path == "/sd/health" || path == "sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		// 请求基础信息
		method := c.Request.Method
		ip := c.ClientIP()

		// 如果不是release模式下, 读取body消息体
		if viper.Get("server.runmode") != "release" {
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			}

			// 将消息体再回传会context.request.body
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			log.Debugf("New request come in, path: %s, Method: %s, clientIP: %s body `%s`", path, method, ip, string(bodyBytes))

		} else {
			log.Debugf("New request come in, path: %s, Method: %s, clientIP: %s", path, method, ip)
		}

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		c.Next()

		// 计算耗时
		latency := time.Since(start)

		code, message := -1, ""

		var response handler.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Errorf(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}

		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 3, ""), path, code, message)
	}
}
