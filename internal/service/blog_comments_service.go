package service

import (
	"hmdp-go/internal/repository"
)

// BlogCommentsService 博客评论服务接口
type BlogCommentsService interface {
}

// blogCommentsService 博客评论服务实现
type blogCommentsService struct {
	repo *repository.Repository
}

// NewBlogCommentsService 创建博客评论服务实例
func NewBlogCommentsService(repo *repository.Repository) BlogCommentsService {
	return &blogCommentsService{repo: repo}
}
