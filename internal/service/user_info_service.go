package service

import (
	"hmdp-go/internal/repository"
)

// UserInfoService 用户信息服务接口
type UserInfoService interface {
}

// userInfoService 用户信息服务实现
type userInfoService struct {
	repo *repository.Repository
}

// NewUserInfoService 创建用户信息服务实例
func NewUserInfoService(repo *repository.Repository) UserInfoService {
	return &userInfoService{repo: repo}
}
