package controller

import (
	"hmdp-go/internal/service"
)

// BlogCommentsController 博客评论控制器接口
type BlogCommentsController interface {
}

// blogCommentsController 博客评论控制器实现
type blogCommentsController struct {
	svc *service.Service
}

// NewBlogCommentsController 创建博客评论控制器实例
func NewBlogCommentsController(svc *service.Service) BlogCommentsController {
	return &blogCommentsController{svc: svc}
}
