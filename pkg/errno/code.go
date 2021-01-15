/**
 * 功能描述: 自定义错误信息code
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package errno

// 错误码定义
// 第1位: 服务级别     1(系统级错误)  2(普通错误)
// 第2-3位: 服务模块   01(用户)
// 第4-5位: 错误码	  01(具体错误代码)
var (
	// 通用错误
	OK                  = &Errno{Code: 0, Message: "OK", CNMessage: "成功"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error.", CNMessage: "内部错误"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct.", CNMessage: "解析Json数据异常"}
	ErrRequestBody      = &Errno{Code: 10003, Message: "The request parameters is invalid.", CNMessage: "请求参数格式不正确"}
	ErrResourceNotFound = &Errno{Code: 10004, Message: "The requested resource does not exist.", CNMessage: "请求的资源不存在"}
	ErrToken            = &Errno{Code: 10005, Message: "Error occurred while signing the JSON web token.", CNMessage: "Token生成失败"}
	ErrJsonParse        = &Errno{Code: 10006, Message: "Error occurred while parse json to struct.", CNMessage: "解析Json数据异常"}
	ErrInitWebsocket    = &Errno{Code: 10007, Message: "Error occurred while upgrade the request to websocket.", CNMessage: "创建websocket会话失败"}
	ErrMessageBroadcast = &Errno{Code: 10008, Message: "Error occurred while broadcast message to room users.", CNMessage: "广播房间消息失败"}
	ErrTokenInvalid     = &Errno{Code: 10009, Message: "Token was invalid.", CNMessage: "Token无效"}

	// 数据库错误
	ErrValidation     = &Errno{Code: 20001, Message: "Validation failed.", CNMessage: "验证失败"}
	ErrDatabaseQuery  = &Errno{Code: 20002, Message: "Database query error.", CNMessage: "服务器查询操作失败,请稍后再试"}
	ErrDatabaseUpdate = &Errno{Code: 20003, Message: "Database update error.", CNMessage: "数据库更新操作失败,请稍后再试"}
	ErrDatabaseDelete = &Errno{Code: 20004, Message: "Database delete error.", CNMessage: "数据库删除操作失败,请稍后再试"}
	ErrDatabaseCreate = &Errno{Code: 20005, Message: "Database create error.", CNMessage: "数据库插入操作失败,请稍后再试"}

	// Redis错误
	ErrRedisSet    = &Errno{Code: 20021, Message: "Redis set cmd error.", CNMessage: "写入缓存数据失败,请稍后再试"}
	ErrRedisDelete = &Errno{Code: 20022, Message: "Redis delete data error.", CNMessage: "删除缓存数据失败,请稍后再试"}
	ErrRedisGet    = &Errno{Code: 20023, Message: "Redis get data error.", CNMessage: "获取缓存数据失败,请稍后再试"}

	// 短信错误
	ErrSMSRegister     = &Errno{Code: 20031, Message: "Sms register info send error.", CNMessage: "发送验证码失败,请稍后再试"}
	ErrSMSCodeExist    = &Errno{Code: 20032, Message: "Sms verification code already send.", CNMessage: "验证码已经发送,请注意查收"}
	ErrSMSCodeInvalid  = &Errno{Code: 20033, Message: "Sms verification code was invalid.", CNMessage: "验证码不正确"}
	ErrSMSCodeNotFound = &Errno{Code: 20034, Message: "Sms verification code was not found.", CNMessage: "验证码不存在,请先获取验证码"}
)
