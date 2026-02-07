package controller

import (
	"hmdp-go/internal/service"
)

// SeckillVoucherController 秒杀优惠券控制器接口
type SeckillVoucherController interface {
}

// seckillVoucherController 秒杀优惠券控制器实现
type seckillVoucherController struct {
	svc *service.Service
}

// NewSeckillVoucherController 创建秒杀优惠券控制器实例
func NewSeckillVoucherController(svc *service.Service) SeckillVoucherController {
	return &seckillVoucherController{svc: svc}
}
