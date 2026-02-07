package repository

import "gorm.io/gorm"

// VoucherRepository 代金券仓库接口
type VoucherRepository interface {
}

// voucherRepository 代金券仓库实现
type voucherRepository struct {
	db *gorm.DB
}

// NewVoucherRepository 创建代金券仓库实例
func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db: db}
}
