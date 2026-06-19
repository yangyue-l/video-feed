package database

import (
	"context"
	"fmt"
	"log"
	"video_feed/config"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

// InitRedis 初始化Redis连接
func InitRedis(cfg *config.Config) error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	if err := RDB.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("连接Redis失败: %w", err)
	}

	log.Println("Redis连接成功")
	return nil
}

// GetRedis 获取Redis客户端
func GetRedis() *redis.Client {
	return RDB
}
