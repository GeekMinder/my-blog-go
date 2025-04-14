package model

import (
	"fmt"
	"os"

	"github.com/GeekMinder/my-blog-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.DBName,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.SSLMode,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数：%v\n", err)
		os.Exit(1)
	}

	// 自动迁移
	err = db.AutoMigrate(
		// 文章
		&Article{},
		// 分类
		&Category{},
		// 登录
		&Auth{},
	)
	if err != nil {
		panic("自动迁移失败: " + err.Error())
	}

	fmt.Println("数据库连接成功")
}
