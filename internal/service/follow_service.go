package service

import (
	"hmdp-go/internal/repository"
)

// FollowService 关注服务接口
type FollowService interface {
}

// followService 关注服务实现
type followService struct {
	repo *repository.Repository
}

// NewFollowService 创建关注服务实例
func NewFollowService(repo *repository.Repository) FollowService {
	return &followService{repo: repo}
}
