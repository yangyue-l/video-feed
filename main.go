package main

import (
	"fmt"
	"log"

	"video_feed/config"
	"video_feed/internal/database"
	"video_feed/internal/logger"
	"video_feed/internal/router"
	"video_feed/internal/utils"
)

// @title           Video Feed API
// @version         1.0
// @description     类似抖音的Feed流视频服务API
func main() {
	// 1. 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 2. 初始化Logger
	if err := logger.Init(&logger.LogConfig{
		Filename:   cfg.Log.Filename,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
		Level:      cfg.Log.Level,
	}, cfg.Server.Mode); err != nil {
		log.Fatalf("初始化Logger失败: %v", err)
	}

	// 3. 初始化数据库连接
	if err := database.InitMySQL(cfg); err != nil {
		log.Fatalf("初始化MySQL失败: %v", err)
	}
	if err := database.InitRedis(cfg); err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}

	// 4. 初始化阿里云OSS（可选）
	if cfg.OSS.Endpoint != "" {
		if err := utils.InitOSS(&cfg.OSS); err != nil {
			log.Printf("初始化OSS失败（可选功能）: %v", err)
		}
	}

	// 5. 初始化Snowflake
	if err := utils.InitSnowflake(); err != nil {
		log.Fatalf("初始化Snowflake失败: %v", err)
	}

	// 6. 设置路由
	r := router.SetupRouter(cfg)

	// 7. 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("服务器启动在 %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
