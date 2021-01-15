/**
 * 功能描述: 数据库初始化
 * @Date: 2019-04-15
 * @author: lixiaoming
 */
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/xm5646/log"
)

type Database struct {
	Instance *gorm.DB
}

var DB *Database

// 根据数据库模式初始化数据库
func (db *Database) Init() {
	if viper.GetString("db.runmode") == "pro" {
		DB = &Database{
			Instance: GetProDB(),
		}
	} else {
		DB = &Database{
			Instance: GetDevDB(),
		}
	}

}

func openDB(username, password, addr, name string) *gorm.DB {
	log.Infof("mysql connection string: %s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	log.Infof("Database connected.")

	setupDB(db)

	return db
}

// 设置数据库连接数和是否打印日志
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("db.gormlog"))
	db.DB().SetMaxOpenConns(viper.GetInt("db.db_max_conn"))
	db.DB().SetMaxIdleConns(viper.GetInt("db.db_max_idle"))
}

func GetDevDB() *gorm.DB {
	return InitDevDB()
}

func InitDevDB() *gorm.DB {
	return openDB(viper.GetString("db.dev.username"),
		viper.GetString("db.dev.password"),
		viper.GetString("db.dev.addr"),
		viper.GetString("db.dev.name"))
}

func GetProDB() *gorm.DB {
	return InitProDB()
}

func InitProDB() *gorm.DB {
	return openDB(viper.GetString("db.pro.username"),
		viper.GetString("db.pro.password"),
		viper.GetString("db.pro.addr"),
		viper.GetString("db.pro.name"))

}

func (db *Database) Close() {
	DB.Instance.Close()
	log.Infof("Database closed.")
}
