package service

import (
	"hmdp-go/internal/repository"
)

// SeckillVoucherService 秒杀优惠券服务接口
type SeckillVoucherService interface {
}

// seckillVoucherService 秒杀优惠券服务实现
type seckillVoucherService struct {
	repo *repository.Repository
}

// NewSeckillVoucherService 创建秒杀优惠券服务实例
func NewSeckillVoucherService(repo *repository.Repository) SeckillVoucherService {
	return &seckillVoucherService{repo: repo}
}
