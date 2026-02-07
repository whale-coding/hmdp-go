package controller

import (
	"hmdp-go/internal/service"
)

// VoucherOrderController 代金券订单控制器接口
type VoucherOrderController interface {
}

// voucherOrderController 代金券订单控制器实现
type voucherOrderController struct {
	svc *service.Service
}

// NewVoucherOrderController 创建代金券订单控制器实例
func NewVoucherOrderController(svc *service.Service) VoucherOrderController {
	return &voucherOrderController{svc: svc}
}
