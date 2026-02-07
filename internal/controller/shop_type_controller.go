package controller

import (
	"hmdp-go/internal/service"
)

// ShopTypeController 店铺类型控制器接口
type ShopTypeController interface {
}

// shopTypeController 店铺类型控制器实现
type shopTypeController struct {
	svc *service.Service
}

// NewShopTypeController 创建店铺类型控制器实例
func NewShopTypeController(svc *service.Service) ShopTypeController {
	return &shopTypeController{svc: svc}
}
