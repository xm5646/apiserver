/**
 * 功能描述: 创建用户
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package user

import (
	. "apiserver/handler"
	"apiserver/model/user"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
	}
	u := user.User{
		Username: r.Username,
		Password: r.Password,
	}

	// 验证数据
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密用户密码
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 插入数据库
	if err := u.Create(); err != nil {
		log.Error("Insert", err, nil)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	resp := CreateResponse{
		Username: r.Username,
	}
	SendResponse(c, nil, resp)
}
