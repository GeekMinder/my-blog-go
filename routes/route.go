package routes

import (
	"github.com/GeekMinder/my-blog-go/api/controller"
	"github.com/GeekMinder/my-blog-go/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 这里是中间件
	r.Use(middleware.Cors())

	// 展示路由
	front := r.Group("/api/v1")
	front.Use(middleware.FrontAuth())
	{
		// 文章展示相关
		article := front.Group("/article")
		{
			article.GET("", controller.GetArticleList)
			article.GET("/:id", controller.GetArticle)
		}

		// 分类展示相关
		category := front.Group("/category")
		{
			category.GET("", controller.GetCategory)
		}

	}

	// 认证
	public := r.Group("/api/admin")
	{
		auth := public.Group("/auth")
		{
			// 注册
			auth.POST("/signup", controller.SignUp)
			// 登录
			auth.POST("/login", controller.Login)
		}
	}

	// 这里是需要鉴权的路由 中台用到的
	private := r.Group("/api/admin")
	private.Use(middleware.JwtAuth())
	{

		article := private.Group("/article")
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

		category := private.Group("/category")
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
	}

	return r
}
