package service

import (
	"hmdp-go/internal/repository"
)

// SignService 签到服务接口
type SignService interface {
}

// signService 签到服务实现
type signService struct {
	repo *repository.Repository
}

// NewSignService 创建签到服务实例
func NewSignService(repo *repository.Repository) SignService {
	return &signService{repo: repo}
}
