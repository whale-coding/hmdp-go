package repository

import "gorm.io/gorm"

// VoucherOrderRepository 代金券订单仓库接口
type VoucherOrderRepository interface {
}

// voucherOrderRepository 代金券订单仓库实现
type voucherOrderRepository struct {
	db *gorm.DB
}

// NewVoucherOrderRepository 创建代金券订单仓库实例
func NewVoucherOrderRepository(db *gorm.DB) VoucherOrderRepository {
	return &voucherOrderRepository{db: db}
}
