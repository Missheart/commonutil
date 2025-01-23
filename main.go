package main

import (
	"fmt"

	lhUtil "github.com/Missheart/commonutil/util"
	"github.com/go-redis/redis/v8"
)

func main() {
	lhUtil.InitRedisConnect(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	lhUtil.HSet("htest", "test111", "test111")
	lhUtil.MyConfig.Db = "htest"
	lhUtil.InitMysqlConnect()
	fmt.Println(lhUtil.MyConfig)
}
