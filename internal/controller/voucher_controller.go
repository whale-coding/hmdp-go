package controller

import (
	"hmdp-go/internal/service"
)

// VoucherController 代金券控制器接口
type VoucherController interface {
}

// voucherController 代金券控制器实现
type voucherController struct {
	svc *service.Service
}

// NewVoucherController 创建代金券控制器实例
func NewVoucherController(svc *service.Service) VoucherController {
	return &voucherController{svc: svc}
}
