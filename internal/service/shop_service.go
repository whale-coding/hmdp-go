package service

import (
	"hmdp-go/internal/repository"
)

// ShopService 商铺服务接口
type ShopService interface {
}

// shopService 商铺服务实现
type shopService struct {
	repo *repository.Repository
}

// NewShopService 创建商铺服务实例
func NewShopService(repo *repository.Repository) ShopService {
	return &shopService{repo: repo}
}
