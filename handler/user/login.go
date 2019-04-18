/**
 * 功能描述: 用户登录
 * @Date: 2019-04-17
 * @author: lixiaoming
 */
package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/model/user"
	"apiserver/pkg/auth"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
)

// @Router /login [post]
func Login(c *gin.Context) {
	// 绑定请求数据到用户User结构体
	var u user.User
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	//根据请求的用户获取用户信息
	d, err := user.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFount, nil)
		return
	}

	// 比较登录请求的密码和数据库中的密码
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// 生成json web token
	t, err := token.Sign(c, token.ContextToken{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})

}
