package repository

import "gorm.io/gorm"

// UserRepository 用户仓库接口
type UserRepository interface {
}

// userRepository 用户仓库实现   UserRepo
type userRepository struct {
	db *gorm.DB
}

// NewUserUserRepository 创建用户仓库实例
func NewUserUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
