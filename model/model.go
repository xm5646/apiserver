/**
 * 功能描述:基础实体类
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package model

import "time"

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeletedAt time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type Token struct {
	Token string `json:"token"`
}
