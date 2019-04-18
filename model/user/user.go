/**
 * 功能描述: Entity 用户
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package user

import (
	"apiserver/model"
	"apiserver/pkg/auth"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	model.BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *User) TableName() string {
	return "tb_users"
}

// 创建用户
func (u *User) Create() error {
	return model.DB.Instance.Create(&u).Error
}

// 删除用户
func DeleteUser(id uint64) error {
	user := User{}
	user.BaseModel.Id = id
	return model.DB.Instance.Delete(&user).Error
}

// 根据用户名获取用户
func GetUser(username string) (*User, error) {
	u := &User{}
	d := model.DB.Instance.Where("username = ?", username).First(&u)
	return u, d.Error
}

// 更新用户cd $GOPATH/src/github.com/swaggo
func (u *User) Update() error {
	return model.DB.Instance.Save(u).Error
}

func (u *User) Count() (int64, error) {
	var count int64
	err := model.DB.Instance.Table(u.TableName()).Count(&count).Error
	return count, err
}

// 加密用户密码
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 字段验证
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
