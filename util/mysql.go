package util

import (
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

var MyswlDb *gorm.DB

var MyConfig *Config = &Config{}

func InitBaseDb() error {
	dsn := MyConfig.User + ":" + MyConfig.Pwd + "@(" + MyConfig.Host + ":" + MyConfig.Port + ")/" + MyConfig.Db + "?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(-1)

	//初始化完成,设置为debug模式
	MyswlDb = db.Debug()

	return nil
}
