package repository

import "gorm.io/gorm"

// BlogCommentsRepository 博客评论仓库接口
type BlogCommentsRepository interface {
}

// blogCommentsRepository 博客评论仓库实现
type blogCommentsRepository struct {
	db *gorm.DB
}

// NewBlogCommentsRepository 创建博客评论仓库实例
func NewBlogCommentsRepository(db *gorm.DB) BlogCommentsRepository {
	return &blogCommentsRepository{db: db}
}
