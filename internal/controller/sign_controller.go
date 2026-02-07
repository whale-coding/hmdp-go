package controller

import (
	"hmdp-go/internal/service"
)

// SignController 签到控制器接口
type SignController interface {
}

// signController 签到控制器实现
type signController struct {
	svc *service.Service
}

// NewSignController 创建签到控制器实例
func NewSignController(svc *service.Service) SignController {
	return &signController{svc: svc}
}
