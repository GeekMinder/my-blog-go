package controller

import (
	"strconv"

	"net/http"

	"github.com/GeekMinder/my-blog-go/model"
	"github.com/GeekMinder/my-blog-go/utils/msg"
	"github.com/gin-gonic/gin"
)

// 查询文章列表 可以通过categoryid查询
func GetArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))

	switch {
	case pageSize >= 10:
		pageSize = 10
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	data, code, total := model.GetArticleList(pageSize, pageNum, uint(categoryId))
	if code == msg.ERROR {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"total":   0,
			"message": msg.GetMsg(code),
			"code":    code,
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"data":    data,
			"total":   total,
			"message": msg.GetMsg(code),
			"code":    code,
		})
	}

}

// 添加文章
func CreateArticle(c *gin.Context) {
	var data model.ArticleCreate
	_ = c.ShouldBindJSON(&data)

	code := model.CreateArticle(&data)
	if code == msg.ERROR {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
		})
	}
}

// 获取单一文章
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticle(uint(id))
	if code == msg.ERROR {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    code,
			"data":    nil,
			"message": msg.GetMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"data":    data,
			"message": msg.GetMsg(code),
		})
	}
}

// 批量删除文章
func DeleteArticle(c *gin.Context) {
	var deleteReqBody struct {
		Ids []uint `json:"ids"`
	}
	if err := c.ShouldBindJSON(&deleteReqBody); err != nil || len(deleteReqBody.Ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    msg.ERROR,
			"message": "无效的请求参数",
		})
		return
	}
	code := model.DeleteArticle(deleteReqBody.Ids)
	if code == msg.ERROR {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
		})
	}
}
