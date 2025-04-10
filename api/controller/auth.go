package controller

import (
	"net/http"

	"github.com/GeekMinder/my-blog-go/model"
	"github.com/GeekMinder/my-blog-go/utils/msg"
	"github.com/gin-gonic/gin"
)

// 注册
func SignUp(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求有误"})
		return
	}

	if requestBody.Username == "" || requestBody.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    msg.ERROR,
			"message": msg.GetMsg(msg.ERROR),
		})
		return
	}

	code := model.SignUp(requestBody.Username, requestBody.Password)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg.GetMsg(code),
	})
}

func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求有误"})
		return
	}
	if loginData.Username == "" || loginData.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    msg.ERROR,
			"message": msg.GetMsg(msg.ERROR),
		})
		return
	}
	code, info, token := model.Login(loginData.Username, loginData.Password)
	if code == msg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
			"data": gin.H{
				"user": gin.H{
					"username": info.Username,
					"role":     info.Role,
				},
				"token": token,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
		})
	}

}
