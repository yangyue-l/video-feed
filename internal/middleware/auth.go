package middleware

import (
	"strings"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 实现JWT认证逻辑
		// 1. 从Header获取token
		// 2. 验证token格式
		// 3. 解析token
		// 4. 将用户信息存入上下文

		// 示例代码框架：
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Unauthorized(c)
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.Unauthorized(c)
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
