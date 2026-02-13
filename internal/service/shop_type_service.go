package service

import (
	"hmdp-go/internal/model"
	"hmdp-go/internal/repository"
)

// ShopTypeService 店铺类型服务接口
type ShopTypeService interface {
	GetShopTypeList() ([]*model.ShopType, error)
}

// shopTypeService 店铺类型服务实现
type shopTypeService struct {
	repo *repository.Repository
}

// NewShopTypeService 创建店铺类型服务实例
func NewShopTypeService(repo *repository.Repository) ShopTypeService {
	return &shopTypeService{repo: repo}
}

// GetShopTypeList 查询店铺类型列表，按sort字段升序
func (s *shopTypeService) GetShopTypeList() ([]*model.ShopType, error) {
	// 调用仓库层查询店铺类型列表，按sort字段升序
	shopTypes, err := s.repo.ShopTypeRepo.FindAllOrderBySort()
	// 添加缓存
	return shopTypes, err
}
