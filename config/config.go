package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// Config 应用配置结构体
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Upload   UploadConfig   `mapstructure:"upload"`
	OSS      OSSConfig      `mapstructure:"oss"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"` // debug, release, test
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireTime int    `mapstructure:"expire_time"` // 过期时间(小时)
}

// UploadConfig 上传配置
type UploadConfig struct {
	MaxSize    int64    `mapstructure:"max_size"`    // 最大文件大小(MB)
	VideoPath  string   `mapstructure:"video_path"`  // 视频存储路径
	CoverPath  string   `mapstructure:"cover_path"`  // 封面存储路径
	AllowTypes []string `mapstructure:"allow_types"` // 允许的视频类型
}

// OSSConfig 阿里云OSS配置
type OSSConfig struct {
	Endpoint        string `mapstructure:"endpoint"`         // OSS端点
	AccessKeyID     string `mapstructure:"access_key_id"`    // AccessKey ID
	AccessKeySecret string `mapstructure:"access_key_secret"` // AccessKey Secret
	BucketName      string `mapstructure:"bucket_name"`      // Bucket名称
	Domain          string `mapstructure:"domain"`           // 访问域名
}

var (
	globalConfig *Config
	once         sync.Once
)

// LoadConfig 加载配置文件
func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AddConfigPath(".")

		// 读取环境变量，将嵌套key用下划线连接
		// 例如: oss.access_key_id -> OSS_ACCESS_KEY_ID
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		if err = viper.ReadInConfig(); err != nil {
			log.Printf("读取配置文件失败: %v", err)
			return
		}

		globalConfig = &Config{}
		if err = viper.Unmarshal(globalConfig); err != nil {
			log.Printf("解析配置文件失败: %v", err)
			return
		}
	})

	return globalConfig, err
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	return globalConfig
}
