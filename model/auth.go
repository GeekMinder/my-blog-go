package model

import (
	"time"

	"github.com/GeekMinder/my-blog-go/utils/msg"
	"github.com/GeekMinder/my-blog-go/utils/pwd"
	"gorm.io/gorm"
)

type Auth struct {
	// id
	ID uint `gorm:"primary_key" json:"id"`
	// 用户名
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	// 密码
	Password string `gorm:"type:varchar(500);not null" json:"password"`
	// 角色 1 管理员 2 普通用户
	Role int `gorm:"type:int;DEFAULT:1" json:"role"`
	// 创建时间
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 更新时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 最后登录时间
	LastLogin time.Time `gorm:"autoUpdateTime" json:"last_login"`
}

// 检查用户的存在性 注册要用
// @username 用户名
// @return bool true 存在 false 不存在
func CheckUserExist(username string) bool {
	var userInfo Auth
	result := db.Where("username =?", username).First(&userInfo)
	if result.Error == gorm.ErrRecordNotFound {
		// 用户名不存在
		return false
	} else if result.Error != nil {
		// 其他错误
		return false
	} else {
		return true
	}
}

// 注册
func SignUp(username string, password string) int {
	// 检查用户名是否存在
	if CheckUserExist(username) {
		return msg.ERROR_USERNAME_USED
	}

	// 创建用户
	myPassword, pwdErr := pwd.EncryptPassword(password)
	if pwdErr != nil {
		return msg.ERROR
	}

	user := Auth{Username: username, Password: myPassword, Role: 1}

	err := db.Create(&user).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 登录
func Login(username string, password string) (int, Auth) {
	var userInfo Auth
	result := db.Where("username =?", username).First(&userInfo)
	if result.Error == gorm.ErrRecordNotFound {
		// 用户名不存在
		return msg.ERROR_USER_NOT_EXIST, Auth{}
	}

	// 验证密码
	if isCorrect := pwd.ValidatePassword(password, userInfo.Password); isCorrect == false {
		return msg.ERROR_PASSWORD_WRONG, Auth{}
	}

	// 更新最后登录时间
	userInfo.LastLogin = time.Now()
	db.Save(&userInfo)

	return msg.SUCCESS, userInfo
}
