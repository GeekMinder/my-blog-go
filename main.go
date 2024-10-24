package main

import (
	"fmt"
	"log"

	"github.com/GeekMinder/my-blog-go/config"
	"github.com/GeekMinder/my-blog-go/model"
	"github.com/GeekMinder/my-blog-go/routes"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatal("加载配置失败:", err)
	}

	// 初始化数据库
	model.InitDB()

	// 引入路由组件
	route := routes.InitRouter()

	// 启动服务器
	serverAddr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	if err := route.Run(serverAddr); err != nil {
		log.Fatal("启动服务器失败:", err)
	}

}
