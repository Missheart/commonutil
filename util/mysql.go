package util

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User string
	Pwd  string
	Db   string
	Host string
	Port string
}

/*
* 全局操作mysql数据库客户端
 */
var MysqlDb *gorm.DB

/*
* mysql连接配置变量
 */
var MyConfig Config

/*
* 初始化mysql连接
 */
func InitMysqlConnect() error {
	log.Println("【Mysql】开始初始化Mysql，检查配置项：", MyConfig)
	dsn := MyConfig.User + ":" + MyConfig.Pwd + "@(" + MyConfig.Host + ":" + MyConfig.Port + ")/" + MyConfig.Db + "?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("【Mysql】开始初始化mysql连接池...")
	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(-1)
	//初始化完成,设置为debug模式
	MysqlDb = db.Debug()
	log.Println("【Mysql】初始化完成...")
	return nil
}
