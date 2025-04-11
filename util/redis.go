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
var RedisClient *redis.Client

/*
* 初始化redis连接
 */
func InitRedisConnect(config *redis.Options) {
	log.Println("【Redis】开始初始化Redis，检查配置项：", config)

	RedisClient = redis.NewClient(config)

	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()

	_, err2 := RedisClient.Ping(ctx).Result()
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
	do, err := RedisClient.Get(ctx, key).Result()
	checkIsRedisNilErr(err)
	return do
}

/*
* 设置redis数据
 */
func Set(key string, value string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := RedisClient.Set(ctx, key, value, 0).Result()
	checkIsRedisNilErr(err)
}

/*
* 删除redis 键
 */
func Del(key string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := RedisClient.Del(ctx, key).Result()
	checkIsRedisNilErr(err)
}

/*
* 获取redis hash数据
 */
func HGet(key string, field string) string {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	do, err := RedisClient.HGet(ctx, key, field).Result()
	checkIsRedisNilErr(err)
	return do
}

/*
* 设置redis hash数据
 */
func HSet(key string, field string, value string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := RedisClient.HSet(ctx, key, field, value).Result()
	checkIsRedisNilErr(err)
}

/*
* 设置redis 过期时间
 */
func Expired(key string, timeout time.Duration) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	_, err := RedisClient.Expire(ctx, key, timeout).Result()
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

/**
* 添加有序集合
 */
func ZAdd(key string, score float64, member string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	err := RedisClient.ZAdd(ctx, key, &redis.Z{Score: score, Member: member}).Err()
	if err != nil {
		panic(err)
	}
}

/*
*
按score倒序获取集合中第几到第几的元素
*/
func ZRevRange(key string, start int64, end int64) []string {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	members, err := RedisClient.ZRevRange(ctx, key, start, end).Result()
	if err != nil {
		panic(err)
	}
	return members
}

// 使用ZIncrBy将成员分数加1
func ZIncrBy(key string, member string) {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	err := RedisClient.ZIncrBy(ctx, key, 1, member).Err()
	if err != nil {
		panic(err)
	}
}

// 获取某个成员的排名（倒序）
func ZRevRank(key string, member string) int64 {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	rank, err := RedisClient.ZRevRank(ctx, key, member).Result()
	if err != nil {
		panic(err)
	}

	return rank
}

// 获取某个元素的分数
func ZScore(key string, member string) float64 {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	score, err := RedisClient.ZScore(ctx, key, member).Result()
	if err != nil {
		panic(err)
	}

	return score
}

// 将一个key的value自增
func IncrBy(key string, value int64) int64 {
	ctx, cancelFunc := getRedisContext()
	defer cancelFunc()
	result, err := RedisClient.IncrBy(ctx, key, value).Result()
	if err != nil {
		panic(err)
	}
	return result
}
