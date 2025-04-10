package main

import (
	"fmt"

	lhUtil "github.com/Missheart/commonutil/util"
	"github.com/go-redis/redis/v8"
)

func main() {
	lhUtil.InitRedisConnect(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	// lhUtil.HSet("htest", "test111", "test111")
	// lhUtil.MyConfig.Db = "htest"
	// lhUtil.InitMysqlConnect()
	// fmt.Println(lhUtil.MyConfig)
	lhUtil.ZAdd("test", 1, "test1")
	lhUtil.ZAdd("test", 3, "test3")
	lhUtil.ZAdd("test", 2, "test2")
	members := lhUtil.ZRevRange("test", 0, -1)
	fmt.Println(members)
	lhUtil.ZIncrBy("test", "test1")
	lhUtil.ZIncrBy("test", "test1")
	lhUtil.ZIncrBy("test", "test1")
	members1 := lhUtil.ZRevRange("test", 0, -1)
	fmt.Println(members1)
}
