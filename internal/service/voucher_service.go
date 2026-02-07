package service

import (
	"hmdp-go/internal/repository"
)

// VoucherService 代金券服务接口
type VoucherService interface {
}

// voucherService 代金券服务实现
type voucherService struct {
	repo *repository.Repository
}

// NewVoucherService 创建代金券服务实例
func NewVoucherService(repo *repository.Repository) VoucherService {
	return &voucherService{repo: repo}
}
