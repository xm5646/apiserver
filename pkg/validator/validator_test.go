package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type User struct {
	Cc    string `json:"cc" validate:"omitempty,cc"`
	Phone string `json:"phone" validate:"required_with=Cc,omitempty,phone"`
	Age   uint   `json:"age" validate:"min=10"`
}

func TestValid(t *testing.T) {
	var user = User{Cc: "1",Phone: "123", Age: 9}

	if errs := vd.Struct(&user); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			fmt.Printf("%s\n", err.Translate(translators["zh"]))
		}
	}

	/*
		cc不是一个可用的国家码
		phone不是一个可用的手机号码
		age最小只能为10
	*/
}