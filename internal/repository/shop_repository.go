package repository

import "gorm.io/gorm"

// ShopRepository 商铺仓库接口
type ShopRepository interface {
}

// shopRepository 商铺仓库实现
type shopRepository struct {
	db *gorm.DB
}

// NewShopRepository 创建商铺仓库实例
func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepository{db: db}
}
