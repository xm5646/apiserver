/**
 * 功能描述: 用户授权认证表
 * @Date: 2020/4/27
 * @author: lixiaoming
 */
package model

import "time"

//create table pcc_user_auth
//(
//id int auto_increment
//primary key,
//user_id int not null comment '用户ID',
//auth_type varchar(100) not null comment '认证方式',
//auth_account varchar(255) not null comment '认证账号',
//password varchar(255) null comment '认证密码',
//open_id varchar(100) null,
//access_token varchar(255) null,
//refresh_token varchar(255) null,
//create_time timestamp null,
//expire_time timestamp null,
//login_time timestamp null,
//constraint tb_user_auth_auth_type_user_id_uindex
//unique (user_id),
//constraint tb_user_auth_tb_user_id_fk
//foreign key (user_id) references pcc_user (id)
//)
//comment '用户登录认证表';

var (
	AuthTypePhone   = "phone"
	AuthTypeWebChat = "webchat"
	AuthTypeGuest   = "guest"
	AuthTypeApple   = "apple"
)

type UserAuthModel struct {
	ID           int       `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;"`
	UserID       int       `json:"user_id" gorm:"column:user_id;not null;"`
	AuthType     string    `json:"auth_type" gorm:"column:auth_type;not null;"`
	AuthAccount  string    `json:"auth_account" gorm:"column:auth_account;"`
	Password     string    `json:"password,omitempty;" gorm:"column:password;"`
	OpenID       string    `json:"open_id" gorm:"column:open_id"`
	AccessToken  string    `json:"access_token" gorm:"column:access_token;"`
	RefreshToken string    `json:"refresh_token" gorm:"column:refresh_token;"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	ExpireTime   time.Time `json:"expire_time" gorm:"column:expire_time;"`
	LoginTime    time.Time `json:"login_time" gorm:"column:login_time"`
}

var UserAuthTableName = "pcc_user_auth"

func (ua *UserAuthModel) TableName() string {
	return UserAuthTableName
}
