package middleware

import (
	"strings"

	"github.com/GeekMinder/my-blog-go/utils/jwt"
	"github.com/GeekMinder/my-blog-go/utils/msg"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			msg.Response(c, msg.ERROR_TOKEN_NOT_EXIST, nil)
			c.Abort()
			return
		}

		// 检查token格式
		parts := strings.SplitN(tokenHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			msg.Response(c, msg.ERROR_TOKEN_TYPE_WRONG, nil)
			c.Abort()
			return
		}

		// 解析token
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			msg.Response(c, msg.ERROR_TOKEN_WRONG, nil)
			c.Abort()
			return
		}

		// 将用户信息保存到上下文
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func FrontAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
