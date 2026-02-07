package repository

import "gorm.io/gorm"

// UserInfoRepository 用户信息仓库接口
type UserInfoRepository interface {
}

// userInfoRepository 用户信息仓库实现
type userInfoRepository struct {
	db *gorm.DB
}

// NewUserInfoRepository 创建用户信息仓库实例
func NewUserInfoRepository(db *gorm.DB) UserInfoRepository {
	return &userInfoRepository{db: db}
}
