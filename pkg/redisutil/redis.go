package redisutil

import (
	"context"
	"encoding/json"
	"hmdp-go/internal/config"
	"log"
	"strconv"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

// RedisClient 全局Redis客户端
var RedisClient *redis.Client
var Ctx = context.Background()

// Redsync 全局分布式锁客户端
var rs *redsync.Redsync

// InitRedis 初始化Redis
func InitRedis() {
	cfg := config.AppConfig.Redis

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     10,              // 连接池大小
		MinIdleConns: 5,               // 最小空闲连接数
		DialTimeout:  5 * time.Second, // 连接超时
		ReadTimeout:  3 * time.Second, // 读超时
		WriteTimeout: 3 * time.Second, // 写超时
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// 初始化 Redsync 分布式锁
	pool := goredis.NewPool(RedisClient)
	rs = redsync.New(pool)

	log.Println("Successfully connected to Redis")
}

// ==================== String 操作 ====================

// Set 设置字符串值（带过期时间，0表示永不过期）
func Set(key string, value interface{}, expire time.Duration) error {
	return RedisClient.Set(Ctx, key, value, expire).Err()
}

// Get 获取字符串值
func Get(key string) (string, error) {
	return RedisClient.Get(Ctx, key).Result()
}

// SetNX 仅当key不存在时设置（用于分布式锁）
func SetNX(key string, value interface{}, expire time.Duration) (bool, error) {
	return RedisClient.SetNX(Ctx, key, value, expire).Result()
}

// Incr 自增
func Incr(key string) (int64, error) {
	return RedisClient.Incr(Ctx, key).Result()
}

// Decr 自减
func Decr(key string) (int64, error) {
	return RedisClient.Decr(Ctx, key).Result()
}

// ==================== 通用操作 ====================

// Del 删除key
func Del(keys ...string) error {
	return RedisClient.Del(Ctx, keys...).Err()
}

// Exists 判断key是否存在
func Exists(key string) (bool, error) {
	n, err := RedisClient.Exists(Ctx, key).Result()
	return n > 0, err
}

// Expire 设置过期时间
func Expire(key string, expire time.Duration) error {
	return RedisClient.Expire(Ctx, key, expire).Err()
}

// TTL 获取剩余过期时间
func TTL(key string) (time.Duration, error) {
	return RedisClient.TTL(Ctx, key).Result()
}

// ==================== Hash 操作 ====================

// HSet 设置hash字段
func HSet(key, field string, value interface{}) error {
	return RedisClient.HSet(Ctx, key, field, value).Err()
}

// HGet 获取hash字段
func HGet(key, field string) (string, error) {
	return RedisClient.HGet(Ctx, key, field).Result()
}

// HGetAll 获取hash所有字段
func HGetAll(key string) (map[string]string, error) {
	return RedisClient.HGetAll(Ctx, key).Result()
}

// HDel 删除hash字段
func HDel(key string, fields ...string) error {
	return RedisClient.HDel(Ctx, key, fields...).Err()
}

// HIncr hash字段自增
func HIncr(key, field string, incr int64) (int64, error) {
	return RedisClient.HIncrBy(Ctx, key, field, incr).Result()
}

// ==================== List 操作 ====================

// LPush 左侧插入
func LPush(key string, values ...interface{}) error {
	return RedisClient.LPush(Ctx, key, values...).Err()
}

// RPush 右侧插入
func RPush(key string, values ...interface{}) error {
	return RedisClient.RPush(Ctx, key, values...).Err()
}

// LPop 左侧弹出
func LPop(key string) (string, error) {
	return RedisClient.LPop(Ctx, key).Result()
}

// LRange 获取列表范围
func LRange(key string, start, stop int64) ([]string, error) {
	return RedisClient.LRange(Ctx, key, start, stop).Result()
}

// ==================== Set 操作 ====================

// SAdd 添加集合成员
func SAdd(key string, members ...interface{}) error {
	return RedisClient.SAdd(Ctx, key, members...).Err()
}

// SMembers 获取集合所有成员
func SMembers(key string) ([]string, error) {
	return RedisClient.SMembers(Ctx, key).Result()
}

// SIsMember 判断是否是集合成员
func SIsMember(key string, member interface{}) (bool, error) {
	return RedisClient.SIsMember(Ctx, key, member).Result()
}

// SRem 移除集合成员
func SRem(key string, members ...interface{}) error {
	return RedisClient.SRem(Ctx, key, members...).Err()
}

// ==================== ZSet 有序集合操作 ====================

// ZAdd 添加有序集合成员
func ZAdd(key string, score float64, member interface{}) error {
	return RedisClient.ZAdd(Ctx, key, redis.Z{Score: score, Member: member}).Err()
}

// ZRange 按排名获取（升序）
func ZRange(key string, start, stop int64) ([]string, error) {
	return RedisClient.ZRange(Ctx, key, start, stop).Result()
}

// ZRevRange 按排名获取（降序）
func ZRevRange(key string, start, stop int64) ([]string, error) {
	return RedisClient.ZRevRange(Ctx, key, start, stop).Result()
}

// ZScore 获取成员分数
func ZScore(key string, member string) (float64, error) {
	return RedisClient.ZScore(Ctx, key, member).Result()
}

// ZRem 移除有序集合成员
func ZRem(key string, members ...interface{}) error {
	return RedisClient.ZRem(Ctx, key, members...).Err()
}

// ==================== JSON 操作（便捷方法）====================

// SetJSON 存储JSON对象
func SetJSON(key string, value interface{}, expire time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return Set(key, data, expire)
}

// ==================== 分布式锁操作 ====================

// AcquireLock 获取分布式锁
// lockKey: 锁的键名
// 返回值: (mutex, error)
func AcquireLock(lockKey string) (*redsync.Mutex, error) {
	// Redsync 的 TryLock 是非阻塞的，Lock 是阻塞的
	// 这里使用 NewMutex 配合 Lock 实现阻塞型锁
	mutex := rs.NewMutex(lockKey)
	if err := mutex.Lock(); err != nil {
		return nil, err
	}
	return mutex, nil
}

// ReleaseLock 释放分布式锁
func ReleaseLock(mutex *redsync.Mutex) error {
	if _, err := mutex.Unlock(); err != nil {
		return err
	}
	return nil
}

// GetJSON 获取JSON对象
func GetJSON(key string, dest interface{}) error {
	data, err := Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}
