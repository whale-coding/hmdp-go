package repository

import "gorm.io/gorm"

// FollowRepository 关注仓库接口
type FollowRepository interface {
}

// followRepository 关注仓库实现
type followRepository struct {
	db *gorm.DB
}

// NewFollowRepository 创建关注仓库实例
func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{db: db}
}
