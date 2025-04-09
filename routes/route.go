package routes

import (
	"github.com/GeekMinder/my-blog-go/api/controller"
	"github.com/GeekMinder/my-blog-go/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 这里是中间件 暂时没有
	r.Use(middleware.Cors())
	// middleware here
	article := r.Group("/api/article")
	{
		// 获取文章列表 不要加/
		article.GET("", controller.GetArticleList)
		// 添加文章
		article.POST("/add", controller.CreateArticle)
		// 获取单一文章
		article.GET("/:id", controller.GetArticle)
		// 删除文章
		article.POST("/delete", controller.DeleteArticle)
	}

	category := r.Group("/api/category")
	{
		// 获取分类列表 不要加/
		category.GET("", controller.GetCategory)
		// 创建分类
		category.POST("/add", controller.CreateCategory)
		// 修改分类
		category.POST("/edit", controller.EditCategory)
		// 删除分类
		category.POST("/delete", controller.DeleteCategory)
	}

	auth := r.Group("/api/auth")
	{
		// 注册
		auth.POST("/signup", controller.SignUp)
		// 登录
		auth.POST("/login", controller.Login)
	}
	return r
}
