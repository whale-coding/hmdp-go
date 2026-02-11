package service

import (
	"hmdp-go/internal/model"
	"hmdp-go/internal/repository"
	"math"
)

// ShopService 商铺服务接口
type ShopService interface {
	QueryShopByType(typeID uint64, page *model.PaginationRequest, x, y float64) ([]*model.Shop, error)
	GetShopById(id uint64) (*model.Shop, error)
}

// shopService 商铺服务实现
type shopService struct {
	repo *repository.Repository
}

// NewShopService 创建商铺服务实例
func NewShopService(repo *repository.Repository) ShopService {
	return &shopService{repo: repo}
}

// QueryShopByType 根据商铺类型分页查询商铺
func (s *shopService) QueryShopByType(typeID uint64, page *model.PaginationRequest, x, y float64) ([]*model.Shop, error) {
	shops, err := s.repo.ShopRepo.FindByTypeIDWithPage(typeID, page.GetOffset(), page.PageSize)
	if err != nil {
		return nil, err
	}

	// 如果传了坐标，计算距离
	if x != 0 && y != 0 {
		for _, shop := range shops {
			shop.Distance = calcDistance(x, y, shop.X, shop.Y)
		}
	}

	return shops, nil
}

// GetShopById 根据ID查询商铺详情
func (s *shopService) GetShopById(id uint64) (*model.Shop, error) {
	return s.repo.ShopRepo.FindByID(id)
}

// calcDistance 计算两点之间的距离（单位：米）使用Haversine公式
func calcDistance(x1, y1, x2, y2 float64) float64 {
	const earthRadius = 6371000 // 地球半径，单位米

	// 转换为弧度
	lat1 := y1 * math.Pi / 180
	lat2 := y2 * math.Pi / 180
	deltaLat := (y2 - y1) * math.Pi / 180
	deltaLon := (x2 - x1) * math.Pi / 180

	// Haversine公式
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
