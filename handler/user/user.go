/**
 * 功能描述: 用户请求响应解析结构体
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package user

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type PhoneLoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Register bool   `json:"register"`
}
