package repository

import (
	"hmdp-go/internal/model"

	"gorm.io/gorm"
)

// UserInfoRepository 用户信息仓库接口
type UserInfoRepository interface {
	FindByUserID(userID uint64) (*model.UserInfo, error)
}

// userInfoRepository 用户信息仓库实现
type userInfoRepository struct {
	db *gorm.DB
}

// NewUserInfoRepository 创建用户信息仓库实例
func NewUserInfoRepository(db *gorm.DB) UserInfoRepository {
	return &userInfoRepository{db: db}
}

// FindByUserID 根据用户ID查找用户信息
func (r *userInfoRepository) FindByUserID(userID uint64) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	err := r.db.Where("user_id = ?", userID).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
