package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.New()
	// 这里是中间件 暂时没有
	// middleware here
	return router
}
