package service

import (
	"hmdp-go/internal/repository"
)

// BlogService 博客服务接口
type BlogService interface {
}

// blogService 博客服务实现
type blogService struct {
	repo *repository.Repository
}

// NewBlogService 创建博客服务实例
func NewBlogService(repo *repository.Repository) BlogService {
	return &blogService{repo: repo}
}
