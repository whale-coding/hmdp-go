package repository

import "gorm.io/gorm"

// SeckillVoucherRepository 秒杀优惠券仓库接口
type SeckillVoucherRepository interface {
}

// seckillVoucherRepository 秒杀优惠券仓库实现
type seckillVoucherRepository struct {
	db *gorm.DB
}

// NewSeckillVoucherRepository 创建秒杀优惠券仓库实例
func NewSeckillVoucherRepository(db *gorm.DB) SeckillVoucherRepository {
	return &seckillVoucherRepository{db: db}
}
