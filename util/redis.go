package util

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

/*
* 全局操作Redis操作客户端
 */
var redisClient *redis.Client

/*
* 初始化redis连接
 */
func InitRedisConnect(config *redis.Options) {
	log.Println("【Redis】开始初始化Redis，检查配置项：", config)

	redisClient = redis.NewClient(config)

	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()

	_, err2 := redisClient.Ping(ctx).Result()
	if err2 != nil {
		panic("redis 连接失败")
	}
	log.Println("【Redis】redis客户端连接成功...")
}

/*
* 获取redis数据
 */
func Get(key string) string {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	do, err := redisClient.Get(ctx, key).Result()
	checkIsRedisNilErr(err)
	return do
}

/*
* 设置redis数据
 */
func Set(key string, value string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := redisClient.Set(ctx, key, value, 0).Result()
	checkIsRedisNilErr(err)
}

/*
* 删除redis 键
 */
func Del(key string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := redisClient.Del(ctx, key).Result()
	checkIsRedisNilErr(err)
}

/*
* 获取redis hash数据
 */
func HGet(key string, field string) string {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	do, err := redisClient.HGet(ctx, key, field).Result()
	checkIsRedisNilErr(err)
	return do
}

/*
* 设置redis hash数据
 */
func HSet(key string, field string, value string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := redisClient.HSet(ctx, key, field, value).Result()
	checkIsRedisNilErr(err)
}

/*
* 设置redis 过期时间
 */
func Expired(key string, timeout time.Duration) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := redisClient.Expire(ctx, key, timeout).Result()
	checkIsRedisNilErr(err)
}

/*
* 判断redis错误是否为nil
 */
func checkIsRedisNilErr(err error) {
	if err != nil && !errors.Is(err, redis.Nil) {
		panic("redis设置过期时间失败")
	}
}

/*
* 获取redis上下文
 */
func getRedisContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
