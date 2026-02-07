package controller

import (
	"hmdp-go/internal/service"
)

// UserInfoController 用户信息控制器接口
type UserInfoController interface {
}

// userInfoController 用户信息控制器实现
type userInfoController struct {
	svc *service.Service
}

// NewUserInfoController 创建用户信息控制器实例
func NewUserInfoController(svc *service.Service) UserInfoController {
	return &userInfoController{svc: svc}
}
