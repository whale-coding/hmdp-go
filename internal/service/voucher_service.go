package service

import (
	"hmdp-go/internal/model"
	"hmdp-go/internal/repository"
)

// VoucherService 代金券服务接口
type VoucherService interface {
	AddCommonVoucher()
	AddSeckillVoucher()
	GetVoucherByShopId(shopId uint64) ([]model.Voucher, error)
}

// voucherService 代金券服务实现
type voucherService struct {
	repo *repository.Repository
}

// NewVoucherService 创建代金券服务实例
func NewVoucherService(repo *repository.Repository) VoucherService {
	return &voucherService{repo: repo}
}

// AddCommonVoucher 添加普通代金券
func (s *voucherService) AddCommonVoucher() {
}

// AddSeckillVoucher 添加秒杀代金券
func (s *voucherService) AddSeckillVoucher() {
}

// GetVoucherByShopId 根据店铺ID查询代金券
func (s *voucherService) GetVoucherByShopId(shopId uint64) ([]model.Voucher, error) {
	return s.repo.VoucherRepo.FindByShopID(shopId)
}
