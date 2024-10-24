package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		DB       string
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	Server struct {
		Port int
	}
}

var AppConfig Config

func LoadConfig() error {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")      // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME")  // 多次调用以添加多个搜索路径

	viper.AutomaticEnv() // 读取匹配的环境变量

	// 设置默认值
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("server.port", 8080)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			return fmt.Errorf("未找到配置文件: %w", err)
		} else {
			// 配置文件被找到，但产生了另外的错误
			return fmt.Errorf("读取配置文件时发生错误: %w", err)
		}
	}

	// 将配置解析为结构体
	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		return fmt.Errorf("无法解析配置到结构体: %w", err)
	}

	return nil
}
