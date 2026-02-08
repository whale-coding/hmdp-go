package service

import (
	"hmdp-go/internal/model"
	"hmdp-go/internal/repository"
)

// BlogService 博客服务接口
type BlogService interface {
	QueryMyBlog(userID uint64, page *model.PaginationRequest) ([]*model.Blog, error)
}

// blogService 博客服务实现
type blogService struct {
	repo *repository.Repository
}

// NewBlogService 创建博客服务实例
func NewBlogService(repo *repository.Repository) BlogService {
	return &blogService{repo: repo}
}

// QueryMyBlog 查询当前登录用户的博客
func (s *blogService) QueryMyBlog(userID uint64, page *model.PaginationRequest) ([]*model.Blog, error) {
	return s.repo.BlogRepo.FindByUserIDWithPage(userID, page.GetOffset(), page.PageSize)
}
