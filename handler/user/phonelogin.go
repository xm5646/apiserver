/**
 * 功能描述: 用户登录
 * @Date: 2019-04-17
 * @author: lixiaoming
 */
package user

import (
	. "apiserver/handler"
	"github.com/gin-gonic/gin"
)

// @Summary 用户通过手机号登录
// @Description 用户通过手机号登录
// @Tags login
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user body user.PhoneLoginRequest true "用户手机登录凭证"
// @Success 200 {object} handler.Response
// @Router /login/phone [post]
func PhoneLogin(c *gin.Context) {
	// 绑定请求数据到用户User结构体

	SendResponse(c, nil, "ok")
}

// @Summary 根据手机号判断用户是否已经注册
// @Description 根据手机号判断用户是否已经注册,如果已注册返回true, 如果未注册返回false,并发送短信验证码
// @Tags login
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param phoneNumber path string true "手机号"
// @Success 200 {object} handler.Response ""
// @Router /login/phone/check/{phoneNumber} [get]
func CheckPhoneIsRegistered(c *gin.Context) {
	phone := c.Param("phoneNumber")

	SendResponse(c, nil, phone)
}
