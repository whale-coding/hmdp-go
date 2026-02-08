package repository

import (
	"hmdp-go/internal/model"

	"gorm.io/gorm"
)

// ShopTypeRepository 店铺类型仓库接口
type ShopTypeRepository interface {
	FindAllOrderBySort() ([]*model.ShopType, error)
}

// shopTypeRepository 店铺类型仓库实现
type shopTypeRepository struct {
	db *gorm.DB
}

// NewShopTypeRepository 创建店铺类型仓库实例
func NewShopTypeRepository(db *gorm.DB) ShopTypeRepository {
	return &shopTypeRepository{db: db}
}

// FindAllOrderBySort 查询所有店铺类型，按sort字段升序
func (r *shopTypeRepository) FindAllOrderBySort() ([]*model.ShopType, error) {
	var shopTypes []*model.ShopType
	err := r.db.Order("sort ASC").Find(&shopTypes).Error
	if err != nil {
		return nil, err
	}
	return shopTypes, nil
}
