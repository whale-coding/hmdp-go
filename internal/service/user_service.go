package service

import (
	"hmdp-go/internal/repository"
)

// UserService 用户服务接口
type UserService interface {
}

// userService 用户服务实现
type userService struct {
	repo *repository.Repository
}

// NewUserService 创建用户服务实例
func NewUserService(repo *repository.Repository) UserService {
	return &userService{repo: repo}
}
