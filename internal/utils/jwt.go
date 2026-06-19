package utils

import (
	"errors"
	"time"
	"video_feed/config"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT声明
type Claims struct {
	UserID   int64  `json:"user_id,string"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(userID int64, username string) (string, error) {
	// 1. 获取JWT配置
	cfg := config.GetConfig()
	// 2. 创建Claims
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWT.ExpireTime) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "video_feed",
		},
	}
	// 3. 生成Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 4. 返回Token字符串
	return token.SignedString([]byte(cfg.JWT.Secret))

}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	// 1. 获取JWT配置
	cfg := config.GetConfig()
	// 2. 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})
	// 3. 验证Token有效性
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	// 4. 返回Claims
	return nil, errors.New("invalid token")

}
