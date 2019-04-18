/**
 * 功能描述: 执行删除用户
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package user

import (
	"apiserver/handler"
	"apiserver/model/user"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := user.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
