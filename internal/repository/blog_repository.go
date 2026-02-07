package repository

import "gorm.io/gorm"

// BlogRepository 博客仓库接口
type BlogRepository interface {
}

// blogRepository 博客仓库实现
type blogRepository struct {
	db *gorm.DB
}

// NewBlogRepository 创建博客仓库实例
func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db: db}
}
