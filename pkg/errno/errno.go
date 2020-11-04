/**
 * 功能描述: 自定义错误信息
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package errno

import "fmt"

type Errno struct {
	Code      int
	Message   string
	CNMessage string
}

func (err *Errno) Error() string {
	return err.Message
}

type Err struct {
	Code      int
	Message   string
	CNMessage string
	Err       error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, CNMessage: errno.CNMessage, Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func DecodeErr(err error) (int, string, string) {
	if err == nil {
		return OK.Code, OK.Message, OK.CNMessage
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message, typed.CNMessage
	case *Errno:
		return typed.Code, typed.Message, typed.CNMessage
	default:
	}

	return InternalServerError.Code, err.Error(), InternalServerError.CNMessage
}
