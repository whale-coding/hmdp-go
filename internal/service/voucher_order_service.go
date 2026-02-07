package service

import (
	"hmdp-go/internal/repository"
)

// VoucherOrderService 代金券订单服务接口
type VoucherOrderService interface {
}

// voucherOrderService 代金券订单服务实现
type voucherOrderService struct {
	repo *repository.Repository
}

// NewVoucherOrderService 创建代金券订单服务实例
func NewVoucherOrderService(repo *repository.Repository) VoucherOrderService {
	return &voucherOrderService{repo: repo}
}
