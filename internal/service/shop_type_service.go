package service

import (
	"hmdp-go/internal/repository"
)

// ShopTypeService 店铺类型服务接口
type ShopTypeService interface {
}

// shopTypeService 店铺类型服务实现
type shopTypeService struct {
	repo *repository.Repository
}

// NewShopTypeService 创建店铺类型服务实例
func NewShopTypeService(repo *repository.Repository) ShopTypeService {
	return &shopTypeService{repo: repo}
}
