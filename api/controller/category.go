package controller

import (
	"net/http"

	"github.com/GeekMinder/my-blog-go/model"
	"github.com/GeekMinder/my-blog-go/utils/msg"
	"github.com/gin-gonic/gin"
)

// CreateCategory 新增分类
func CreateCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	isOk, _ := model.CheckCategoryByName(data.Name)
	switch isOk {
	// 不存在就创建
	case msg.ERROR_CATEGORY_NOT_EXIST:
		code := model.CreateCategory(&data)
		if code == msg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": msg.GetMsg(msg.SUCCESS),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    code,
				"message": msg.GetMsg(msg.ERROR),
			})
		}

		// 存在就返回已存在
	case msg.ERROR_CATEGORY_EXIST:
		c.JSON(http.StatusOK, gin.H{
			"code":    isOk,
			"message": msg.GetMsg(msg.ERROR_CATEGORY_EXIST),
		})

		// 其他错误 就返回操作失败
	case msg.ERROR:
		c.JSON(http.StatusOK, gin.H{
			"code":    isOk,
			"message": msg.GetMsg(msg.ERROR),
		})
	}
}

// 修改分类的名称 应该只有名称可以改了
func EditCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	isOk, _ := model.CheckCategoryById(data.ID)
	// 如果找到了 就改
	if isOk == msg.SUCCESS {
		code := model.EditCategory(data.ID, data.Name)
		if code == msg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "修改" + msg.GetMsg(msg.SUCCESS),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    code,
				"message": "修改" + msg.GetMsg(msg.ERROR),
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    isOk,
			"data":    "修改" + msg.GetMsg(msg.ERROR),
			"message": "修改" + msg.GetMsg(msg.ERROR),
		})
	}
}

// 删除category
func DeleteCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	isOk, _ := model.CheckCategoryById(data.ID)
	// 如果找到了 就删
	if isOk == msg.SUCCESS {
		code := model.DeleteCategory(data.ID)
		if code == msg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "删除" + msg.GetMsg(msg.SUCCESS),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    code,
				"message": "删除" + msg.GetMsg(msg.ERROR),
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    isOk,
			"data":    "删除" + msg.GetMsg(msg.ERROR),
			"message": "删除" + msg.GetMsg(msg.ERROR),
		})
	}
}

// 获取所有分类
func GetCategory(c *gin.Context) {
	data, code, total := model.GetCategory()
	if code == msg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"data":    data,
			"total":   total,
			"message": msg.GetMsg(code),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    code,
			"data":    nil,
			"total":   0,
			"message": msg.GetMsg(code),
		})
	}
}
