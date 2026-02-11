package repository

import (
	"hmdp-go/internal/model"

	"gorm.io/gorm"
)

// ShopRepository 商铺仓库接口
type ShopRepository interface {
	FindByTypeIDWithPage(typeID uint64, offset, limit int) ([]*model.Shop, error)
	FindByID(id uint64) (*model.Shop, error)
}

// shopRepository 商铺仓库实现
type shopRepository struct {
	db *gorm.DB
}

// NewShopRepository 创建商铺仓库实例
func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepository{db: db}
}

// FindByTypeIDWithPage 根据商铺类型ID分页查询商铺
func (r *shopRepository) FindByTypeIDWithPage(typeID uint64, offset, limit int) ([]*model.Shop, error) {
	var shops []*model.Shop
	err := r.db.Where("type_id = ?", typeID).
		Order("id ASC").
		Offset(offset).
		Limit(limit).
		Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return shops, nil
}

// FindByID 根据ID查询商铺
func (r *shopRepository) FindByID(id uint64) (*model.Shop, error) {
	var shop model.Shop
	err := r.db.Where("id = ?", id).First(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}
