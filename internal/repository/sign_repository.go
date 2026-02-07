package repository

import "gorm.io/gorm"

// SignRepository 签到仓库接口
type SignRepository interface {
}

// signRepository 签到仓库实现
type signRepository struct {
	db *gorm.DB
}

// NewSignRepository 创建签到仓库实例
func NewSignRepository(db *gorm.DB) SignRepository {
	return &signRepository{db: db}
}
