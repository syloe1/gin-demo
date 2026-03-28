package middleware

import (
	"net/http"
	"strings"

	"gin-demo/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1️⃣ 获取 header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未登录",
			})
			c.Abort()
			return
		}

		// 2️⃣ 解析 Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token格式错误",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3️⃣ 校验 token
		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token无效",
			})
			c.Abort()
			return
		}

		// 4️⃣ 存入上下文（后面可以用）
		c.Set("user_id", userID)

		c.Next()
	}
}