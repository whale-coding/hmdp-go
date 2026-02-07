package controller

import (
	"hmdp-go/internal/service"
)

// FollowController 关注控制器接口
type FollowController interface {
}

// followController 关注控制器实现
type followController struct {
	svc *service.Service
}

// NewFollowController 创建关注控制器实例
func NewFollowController(svc *service.Service) FollowController {
	return &followController{svc: svc}
}
