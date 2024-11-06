package model

import (
	"time"

	"github.com/GeekMinder/my-blog-go/utils/msg"
	"gorm.io/gorm"
)

type Category struct {
	// id
	ID uint `gorm:"primary_key" json:"id"`
	// 创建时间
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 更新时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 分类名称
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

type CategoryBasic struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
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

// DeleteCate 删除分类 单个删除
func DeleteCategory(id uint) int {
	// 开启事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 检查分类是否存在
	var count int64
	if err := tx.Model(&Category{}).Where("id = ?", id).Count(&count).Error; err != nil || count == 0 {
		tx.Rollback()
		return msg.ERROR
	}

	// 先删除关联表中的数据
	if err := tx.Table("article_categories").Where("category_id = ?", id).Delete(nil).Error; err != nil {
		tx.Rollback()
		return msg.ERROR
	}

	// 需要硬删除

	if err := tx.Where("id = ? ", id).Unscoped().Delete(&Category{}).Error; err != nil {
		tx.Rollback()
		return msg.ERROR
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 获取所有分类
func GetCategory() ([]CategoryBasic, int, int64) {
	var category []CategoryBasic
	var total int64
	err = db.Model(&Category{}).Find(&category).Count(&total).Error
	if err != nil {
		return nil, msg.ERROR, 0
	}
	return category, msg.SUCCESS, total
}
