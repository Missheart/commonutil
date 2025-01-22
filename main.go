package main

import (
	lhUtil "github.com/Missheart/commonutil/util"
	"github.com/go-redis/redis/v8"
)

func main() {
	lhUtil.InitRedis(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	lhUtil.HSet("htest", "test111", "test111")
}
