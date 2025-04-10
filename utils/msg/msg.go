package msg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// common
	SUCCESS = 200
	ERROR   = 500

	// 文章
	ERROR_ARTICLE_NOT_EXIST = 2001

	// 分类
	ERROR_CATEGORY_EXIST     = 3001
	ERROR_CATEGORY_NOT_EXIST = 3002

	// 登录
	ERROR_USER_NOT_EXIST = 4011
	ERROR_PASSWORD_WRONG = 4012
	ERROR_USERNAME_USED  = 4013

	// token
	ERROR_TOKEN_NOT_EXIST  = 1001
	ERROR_TOKEN_TYPE_WRONG = 1002
	ERROR_TOKEN_WRONG      = 1003
)

var codeMsg = map[int]string{
	// common
	SUCCESS: "操作成功",
	ERROR:   "操作失败",

	// 文章
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",

	// 分类
	ERROR_CATEGORY_EXIST:     "该分类已存在",
	ERROR_CATEGORY_NOT_EXIST: "该分类不存在",

	// 登录
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USERNAME_USED:  "用户名已存在",

	// token
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
}

func GetMsg(code int) string {
	return codeMsg[code]
}

// Response 统一返回格式
func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": GetMsg(code),
		"data":    data,
	})
}
