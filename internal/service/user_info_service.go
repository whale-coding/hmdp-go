package service

import (
	"errors"
	"hmdp-go/internal/model"
	"hmdp-go/internal/repository"
	"log"

	"gorm.io/gorm"
)

// UserInfoService 用户信息服务接口
type UserInfoService interface {
	GetUserInfo(userID uint64) (*model.UserInfo, error)
}

// userInfoService 用户信息服务实现
type userInfoService struct {
	repo *repository.Repository
}

// NewUserInfoService 创建用户信息服务实例
func NewUserInfoService(repo *repository.Repository) UserInfoService {
	return &userInfoService{repo: repo}
}

// GetUserInfo 根据用户ID获取用户信息
func (s *userInfoService) GetUserInfo(userID uint64) (*model.UserInfo, error) {
	userInfo, err := s.repo.UserInfoRepo.FindByUserID(userID)
	if err != nil {
		log.Printf("查询用户信息失败,用户ID: %d,错误: %v", userID, err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, errors.New("查询用户信息失败")
	}
	return userInfo, nil
}
