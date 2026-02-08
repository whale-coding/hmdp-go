package repository

import (
	"hmdp-go/internal/model"

	"gorm.io/gorm"
)

// BlogRepository 博客仓库接口
type BlogRepository interface {
	FindByUserIDWithPage(userID uint64, offset, limit int) ([]*model.Blog, error)
}

// blogRepository 博客仓库实现
type blogRepository struct {
	db *gorm.DB
}

// NewBlogRepository 创建博客仓库实例
func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db: db}
}

// FindByUserIDWithPage 根据用户ID分页查询博客
func (r *blogRepository) FindByUserIDWithPage(userID uint64, offset, limit int) ([]*model.Blog, error) {
	var blogs []*model.Blog
	err := r.db.Where("user_id = ?", userID).
		Order("create_time DESC").
		Offset(offset).
		Limit(limit).
		Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}
