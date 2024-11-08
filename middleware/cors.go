package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 添加错误处理和日志
		defer func() {
			if err := recover(); err != nil {
				log.Printf("CORS middleware error: %v", err)
			}
		}()
		method := c.Request.Method

		//接收客户端发送的origin （重要！）
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		//服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//允许跨域设置可以返回其他子段，可以自定义字段
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session, Content-Type")
		// 允许浏览器（客户端）可以解析的头部 （重要）
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Token")
		//设置缓存时间
		c.Header("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		c.Header("Access-Control-Allow-Credentials", "true")

		//允许类型校验
		if method == "OPTIONS" {
			// c.AbortWithStatus(http.StatusNoContent)
			c.Status(http.StatusOK)
			c.Abort()
			return
		}
		c.Next()
	}
}
