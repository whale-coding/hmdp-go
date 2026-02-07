package redisutil

import (
	"context"
	"hmdp-go/internal/config"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// Rdb 全局Redis客户端,通过redis.Rdb.xxxx 使用
var Rdb *redis.Client
var Ctx = context.Background()

// InitRedis 初始化Redis
func InitRedis() {
	cfg := config.AppConfig.Redis

	// 关键点：初始化redisClient
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试是否能够连通
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	log.Println("Successfully connected to Redis")

	// 赋值全局变量
	Rdb = RedisClient
}
