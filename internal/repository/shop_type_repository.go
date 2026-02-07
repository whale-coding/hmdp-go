package repository

import "gorm.io/gorm"

// ShopTypeRepository 店铺类型仓库接口
type ShopTypeRepository interface {
}

// shopTypeRepository 店铺类型仓库实现
type shopTypeRepository struct {
	db *gorm.DB
}

// NewShopTypeRepository 创建店铺类型仓库实例
func NewShopTypeRepository(db *gorm.DB) ShopTypeRepository {
	return &shopTypeRepository{db: db}
}
