package util

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis(config *redis.Options) {
	redisClient = redis.NewClient(config)

	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()

	_, err2 := redisClient.Ping(ctx).Result()
	if err2 != nil {
		panic("redis 连接失败")
	}
}

func Get(key string) string {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	do, err := redisClient.Get(ctx, key).Result()
	if !checkIsRedisErr(err) {
		panic("redis获取数据失败")
	}
	return do
}

func Set(key string, value string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := redisClient.Set(ctx, key, value, 0).Result()
	if !checkIsRedisErr(err) {
		panic("redis设置数据失败")
	}
}

func Del(key string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := redisClient.Del(ctx, key).Result()
	if !checkIsRedisErr(err) {
		panic("redis删除数据失败")
	}
}

func HGet(key string, field string) string {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	do, err := redisClient.HGet(ctx, key, field).Result()
	if !checkIsRedisErr(err) {
		panic("redis获取数据失败")
	}
	return do
}

func HSet(key string, field string, value string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := redisClient.HSet(ctx, key, field, value).Result()
	if !checkIsRedisErr(err) {
		panic("redis设置数据失败")
	}
}

func Expired(key string, timeout time.Duration) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := redisClient.Expire(ctx, key, timeout).Result()
	if !checkIsRedisErr(err) {
		panic("redis设置过期时间失败")
	}
}

func checkIsRedisErr(err error) bool {
	if err != nil && !errors.Is(err, redis.Nil) {
		return false
	}
	return true
}

func getRedisContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
