package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"zixu.com/cmdAdmin/util"
)

// DB 数据库链接单例
var DB *gorm.DB
var err error

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	DB, err = gorm.Open("mysql", connString)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//DB.LogMode(true)
	//DB = &(*db)
	//设置连接池
	//空闲
	DB.DB().SetMaxIdleConns(50)
	//打开
	DB.DB().SetMaxOpenConns(100)
	//超时
	DB.DB().SetConnMaxLifetime(time.Second * 30)
	migration()

}
func GetDB() *gorm.DB {
	return DB
}
