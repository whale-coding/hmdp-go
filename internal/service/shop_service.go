package service

import (
	"fmt"
	"hmdp-go/internal/constant"
	"hmdp-go/internal/model"
	"hmdp-go/internal/repository"
	"hmdp-go/pkg/redisutil"
	"hmdp-go/pkg/util"
	"log"
	"math"
)

// ShopService 商铺服务接口
type ShopService interface {
	QueryShopByType(typeID uint64, page *model.PaginationRequest, x, y float64) ([]*model.Shop, error)
	GetShopById(id uint64) (*model.Shop, error)
	CreateShop(shop *model.Shop) error
	UpdateShop(shop *model.Shop) error
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
	// 商户详情缓存key
	shopCacheKey := constant.CACHE_SHOP_KEY + util.Uint64ToString(id)
	// 1.从redis中查询商铺缓存
	shopStr, err := redisutil.Get(shopCacheKey)
	// 2.判断缓存中是否有商铺信息
	if err == nil && shopStr != "" {
		// 3.有，反序列化后返回
		var shop model.Shop
		if err := util.Unmarshal(shopStr, &shop); err != nil {
			return nil, err // 反序列化失败则返回错误
		}
		log.Println("shop found in cache")
		return &shop, nil
	}
	// 判断命中的是否是空值（缓存穿透）
	if err == nil && shopStr == "" {
		log.Println("shop not found in cache, but empty value exists")
		return nil, fmt.Errorf("shop not found")
	}

	// 4、缓存中没有，使用分布式锁防止缓存击穿（避免热点数据失效时多个请求并发查询数据库）
	lockKey := constant.LOCK_SHOP_KEY + util.Uint64ToString(id)
	mutex, err := redisutil.AcquireLock(lockKey)
	if err != nil {
		log.Println("failed to acquire lock:", err)
		return nil, err
	}
	// 确保锁被释放
	defer func() {
		if err := redisutil.ReleaseLock(mutex); err != nil {
			log.Println("failed to release lock:", err)
		}
	}()

	// 再次检查缓存（double check），防止其他线程已经更新过缓存
	shopStr, err = redisutil.Get(shopCacheKey)
	if err == nil && shopStr != "" {
		var shop model.Shop
		if err := util.Unmarshal(shopStr, &shop); err != nil {
			return nil, err
		}
		log.Println("shop found in cache after acquiring lock")
		return &shop, nil
	}

	// 缓存中确实没有，查询数据库
	shop, err := s.repo.ShopRepo.FindByID(id)
	// 4.1 数据库中没有，返回错误（数据库查询不到数据时，err != nil）
	if err != nil {
		// 将空值存入redis缓存，设置过期时间，避免缓存穿透
		if err := redisutil.Set(shopCacheKey, "", constant.CACHE_NULL_TTL); err != nil {
			log.Println("failed to set null cache:", err)
		}
		log.Println("shop not found")
		return nil, fmt.Errorf("shop not found")
	}
	// 4.2 数据库中有，返回商铺信息并将其存入redis缓存
	// 5.将商铺信息序列化后存入redis缓存
	shopStr, err = util.Marshal(shop)
	if err != nil {
		return nil, err // 序列化失败则返回错误
	}
	// 5.1 存入redis缓存，设置过期时间
	if err := redisutil.Set(shopCacheKey, shopStr, constant.CACHE_SHOP_TTL); err != nil {
		log.Println("failed to set shop cache:", err)
	}
	log.Println("shop found in database and cached")

	// 6.返回商铺信息
	return shop, nil
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

// CreateShop 创建商铺
func (s *shopService) CreateShop(shop *model.Shop) error {
	return s.repo.ShopRepo.CreateShop(shop)
}

// 更新商铺信息（记得加事务）
func (s *shopService) UpdateShop(shop *model.Shop) error {
	// 店铺id不能为空
	if shop.ID == 0 {
		return fmt.Errorf("shop id cannot be empty")
	}
	// 1.更新数据库
	if err := s.repo.ShopRepo.UpdateShop(shop); err != nil {
		return err
	}
	// 2.删除缓存
	shopCacheKey := constant.CACHE_SHOP_KEY + util.Uint64ToString(shop.ID)
	if err := redisutil.Del(shopCacheKey); err != nil {
		return err
	}
	return nil
}
