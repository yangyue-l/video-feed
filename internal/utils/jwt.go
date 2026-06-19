package utils

import (
	"errors"
	"time"
	"video_feed/config"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT声明
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(userID uint, username string) (string, error) {
	// TODO: 实现JWT Token生成
	// 1. 获取JWT配置
	// 2. 创建Claims
	// 3. 生成Token
	// 4. 返回Token字符串

	// 示例代码框架：
	cfg := config.GetConfig()

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWT.ExpireTime) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "video_feed",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	// TODO: 实现JWT Token解析
	// 1. 获取JWT配置
	// 2. 解析Token
	// 3. 验证Token有效性
	// 4. 返回Claims

	// 示例代码框架：
	cfg := config.GetConfig()

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
