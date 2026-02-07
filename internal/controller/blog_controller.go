package controller

import (
	"hmdp-go/internal/service"
)

// BlogController 博客控制器接口
type BlogController interface {
}

// blogController 博客控制器实现
type blogController struct {
	svc *service.Service
}

// NewBlogController 创建博客控制器实例
func NewBlogController(svc *service.Service) BlogController {
	return &blogController{svc: svc}
}
