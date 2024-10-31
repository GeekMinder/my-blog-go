package model

import (
	"github.com/GeekMinder/my-blog-go/utils/msg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
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
func GetArticleList(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Preload("Categories", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	}).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Order("created_at DESC").
		Find(&articleList).Error
	db.Model(&Article{}).Count(&total)
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
