package repository

import (
	"hmdp-go/internal/model"

	"gorm.io/gorm"
)

// UserRepository 用户仓库接口
type UserRepository interface {
	FindByPhone(phone string) (*model.User, error)
	FindByID(id uint64) (*model.User, error)
	Create(user *model.User) error
}

// userRepository 用户仓库实现
type userRepository struct {
	db *gorm.DB
}

// NewUserUserRepository 创建用户仓库实例
func NewUserUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// FindByPhone 根据手机号查找用户
func (r *userRepository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID 根据ID查找用户
func (r *userRepository) FindByID(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create 创建用户
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}
