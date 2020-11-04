/**
 * 功能描述: 用户表
 * @Date: 2020/4/26
 * @author: lixiaoming
 */
package model

import (
	"apiserver/pkg/errno"
)

//create table pcc_user
//(
//id int auto_increment comment '用户ID'
//primary key,
//nickname varchar(100) null comment '用户昵称',
//avatar varchar(2000) null comment '头像URL',
//sex tinyint null comment '性别',
//age int null comment '年龄',
//sign varchar(200) null comment '个性签名',
//company varchar(200) null comment '公司名称',
//business int null comment '行业',
//phone varchar(20) null comment '手机号'
//)
//comment '用户表' charset=latin1;

var (
	AvatarDefault = "http://mobile.mythvip.top/imasges/logo.png"
)

type UserModel struct {
	ID       int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;"`
	NickName string `json:"nick_name" gorm:"column:nickname"`
	Avatar   string `json:"avatar" gorm:"column:avatar;default:'user_avatar.png'"`
	Sex  uint8 `json:"sex" gorm:"column:sex;"`
	Age uint8 `json:"age" gorm:"column:age;"`
	Sign string `json:"sign" gorm:"column:sign;"`
	Company string `json:"company" gorm:"column:company;"`
	Business uint8 `json:"business" gorm:"column:business;"`
	Phone string `json:"phone" gorm:"column:phone;"`
}

var UserTableName = "pcc_user"

func (u *UserModel) TableName() string {
	return UserTableName
}

func (u *UserModel) Create(auth *UserAuthModel) *errno.Err {
	tx := DB.Instance.Begin()

	err := tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		return errno.New(errno.ErrDatabaseUpdate, err)

	} else {
		auth.UserID = u.ID
		err = tx.Create(&auth).Error
		if err != nil {
			tx.Rollback()
			return errno.New(errno.ErrDatabaseCreate, err)
		}
		tx.Commit()
		return nil
	}
}

func (u *UserModel) Update() *errno.Err {
	err := DB.Instance.Save(u).Error
	if err != nil {
		return errno.New(errno.ErrDatabaseUpdate, err)
	} else {
		return nil
	}

}
