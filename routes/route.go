package routes

import (
	"github.com/GeekMinder/my-blog-go/api/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 这里是中间件 暂时没有
	// middleware here
	article := r.Group("/api/article")
	{
		// 获取文章列表
		article.GET("/", controller.GetArticleList)
		// 添加文章
		article.POST("/add", controller.CreateArticle)
	}

	category := r.Group("/api/category")
	{
		// 获取分类列表

		// 创建分类
		category.POST("/add", controller.CreateCategory)
		// 修改分类
		category.POST("/edit", controller.EditCategory)
		// 删除分类
		category.POST("delete", controller.DeleteCategory)
	}

	return r
}
