package model

import (
	"github.com/GeekMinder/my-blog-go/utils/msg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CreateCate 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

// CheckCategory 查询分类是否存在
// @name 传过来的name字符串
func CheckCategoryByName(name string) (code int, category Category) {
	result := db.Select("id").Where("name = ?", name).First(&category)
	if result.Error == gorm.ErrRecordNotFound {
		// 分类不存在
		return msg.ERROR_CATEGORY_NOT_EXIST, Category{}
	} else if result.Error != nil {
		// 其他错误
		return msg.ERROR, Category{}
	}
	// 分类存在
	return msg.ERROR_CATEGORY_EXIST, category
}

// @id 传过来的name字符串
func CheckCategoryById(id uint) (code int, category Category) {
	result := db.Select("id").Where("id = ?", id).First(&category)
	if result.Error == gorm.ErrRecordNotFound {
		// 分类不存在
		return msg.ERROR_CATEGORY_NOT_EXIST, Category{}
	} else if result.Error != nil {
		// 其他错误
		return msg.ERROR, Category{}
	}
	// 分类存在
	return msg.SUCCESS, category
}

// EditCate 编辑分类信息
func EditCategory(id uint, name string) int {
	var category Category
	err = db.Model(&category).Where("id = ? ", id).Update("name", name).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

// DeleteCate 删除分类
func DeleteCategory(id uint) int {
	var category Category
	// 需要硬删除
	err = db.Where("id = ? ", id).Unscoped().Delete(&category).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 获取所有分类
func GetCategory() ([]Category, int, int64) {
	var category []Category
	var total int64
	err = db.Find(&category).Count(&total).Error
	if err != nil {
		return nil, msg.ERROR, 0
	}
	return category, msg.SUCCESS, total
}
