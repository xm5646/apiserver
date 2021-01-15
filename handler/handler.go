/**
 * 功能描述: API返回入口函数
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package handler

import (
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	CNMessage string      `json:"cn_message"`
	Data      interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message, cnMessage := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:      code,
		Message:   message,
		CNMessage: cnMessage,
		Data:      data,
	})
}
