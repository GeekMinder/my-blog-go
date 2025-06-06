package model

import (
	"time"

	"github.com/GeekMinder/my-blog-go/utils/msg"
	"gorm.io/gorm"
)

type Article struct {
	// id
	ID uint `gorm:"primary_key" json:"id"`
	// 创建时间
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 更新时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 标题
	Title string `json:"title" gorm:"type:varchar(100);not null"`
	// 内容
	Content string `json:"content" gorm:"type:text;not null"`
	// 描述
	Desc string `json:"desc" gorm:"type:varchar(200);not null"`
	// 阅读量
	ViewCount int64 `json:"view_count" gorm:"type:bigint;default:0"`
	// 点赞数
	LikeCount int64 `json:"like_count" gorm:"type:bigint;default:0"`
	// 评论数
	CommentCount int64 `json:"comment_count" gorm:"type:bigint;default:0"`
	// 分类
	Categories []Category `json:"categories" gorm:"many2many:article_categories"`
}

type ArticleCreate struct {
	Title         string   `json:"title" binding:"required"`
	Content       string   `json:"content" binding:"required"`
	Desc          string   `json:"desc" binding:"required"`
	CategoryNames []string `json:"categories" binding:"required"`
}

// 获取文章列表
func GetArticleList(pageSize int, pageNum int, categoryId uint) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64

	query := db.Preload("Categories", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	})

	// 如果提供了分类ID，添加分类过滤条件
	if categoryId > 0 {
		query = query.Joins("JOIN article_categories ON articles.id = article_categories.article_id").
			Where("article_categories.category_id = ?", categoryId)
	}

	err = query.
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Order("created_at DESC").
		Find(&articleList).Error

	// 计算总数时也需要考虑分类过滤
	countQuery := db.Model(&Article{})
	if categoryId > 0 {
		countQuery = countQuery.Joins("JOIN article_categories ON articles.id = article_categories.article_id").
			Where("article_categories.category_id = ?", categoryId)
	}
	countQuery.Count(&total)

	if err != nil {
		return nil, msg.ERROR, 0
	}
	return articleList, msg.SUCCESS, total
}

// 增加文章
func CreateArticle(data *ArticleCreate) int {
	// 批量查询已存在的分类
	var existingCategories []Category
	if err := db.Where("name IN ?", data.CategoryNames).Find(&existingCategories).Error; err != nil {
		return msg.ERROR
	}
	// 找出需要新建的分类名称
	existingNames := make(map[string]bool)
	for _, cat := range existingCategories {
		existingNames[cat.Name] = true
	}

	var newCategories []Category
	for _, name := range data.CategoryNames {
		if !existingNames[name] {
			newCategories = append(newCategories, Category{Name: name})
		}
	}

	// 批量创建新分类（如果有的话）
	if len(newCategories) > 0 {
		if err := db.Create(&newCategories).Error; err != nil {
			return msg.ERROR
		}
		existingCategories = append(existingCategories, newCategories...)
	}

	// 创建文章并关联分类
	article := Article{
		Title:      data.Title,
		Content:    data.Content,
		Desc:       data.Desc,
		Categories: existingCategories,
	}

	if err := db.Create(&article).Error; err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 获取单一文章
func GetArticle(id uint) (Article, int) {
	var article Article
	if err := db.Preload("Categories").First(&article, id).Error; err != nil {
		return article, msg.ERROR
	}
	return article, msg.SUCCESS
}

// 删除文章
func DeleteArticle(ids []uint) int {
	// 开启事务
	tx := db.Begin()
	// 确保事务最后一定会提交或回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 检查是否有这些文章存在
	var count int64
	if err := tx.Model(&Article{}).Where("id IN ?", ids).Count(&count).Error; err != nil || count == 0 {
		tx.Rollback()
		return msg.ERROR
	}

	// 先删除关联表中的数据
	if err := tx.Table("article_categories").Where("article_id IN ?", ids).Delete(nil).Error; err != nil {
		tx.Rollback()
		return msg.ERROR
	}
	// 删除文章
	if err := tx.Where("id In ?", ids).Unscoped().Delete(&Article{}).Error; err != nil {
		tx.Rollback()
		return msg.ERROR
	}
	// 所有操作都成功，提交事务
	if err := tx.Commit().Error; err != nil {
		return msg.ERROR
	}

	return msg.SUCCESS
}
