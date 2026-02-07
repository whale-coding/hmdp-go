package controller

import (
	"hmdp-go/internal/service"
)

// UserController 用户控制器接口
type UserController interface {
}

// userController 用户控制器实现
type userController struct {
	svc *service.Service
}

// NewUserController 创建用户控制器实例
func NewUserController(svc *service.Service) UserController {
	return &userController{svc: svc}
}
