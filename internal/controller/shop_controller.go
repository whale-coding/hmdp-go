package controller

import (
	"hmdp-go/internal/service"
)

// ShopController 商铺控制器接口
type ShopController interface {
}

// shopController 商铺控制器实现
type shopController struct {
	svc *service.Service
}

// NewShopController 创建商铺控制器实例
func NewShopController(svc *service.Service) ShopController {
	return &shopController{svc: svc}
}
