package controller

import (
	"strconv"

	"net/http"

	"github.com/GeekMinder/my-blog-go/model"
	"github.com/GeekMinder/my-blog-go/utils/msg"
	"github.com/gin-gonic/gin"
)

// 查询文章列表
func GetArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	switch {
	case pageSize >= 10:
		pageSize = 10
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	data, code, total := model.GetArticleList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"total":   total,
		"message": msg.GetMsg(code),
		"code":    code,
	})

}

// 添加文章
func CreateArticle(c *gin.Context) {
	var data model.ArticleCreate
	_ = c.ShouldBindJSON(&data)

	code := model.CreateArticle(&data)
	if code == msg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": msg.GetMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": msg.GetMsg(code),
		})
	}

}
