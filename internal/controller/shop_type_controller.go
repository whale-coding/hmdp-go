package controller

import (
	"hmdp-go/internal/constant"
	"hmdp-go/internal/service"
	"hmdp-go/pkg/result"

	"github.com/gin-gonic/gin"
)

// ShopTypeController 店铺类型控制器接口
type ShopTypeController interface {
	GetShopTypeList(ctx *gin.Context)
}

// shopTypeController 店铺类型控制器实现
type shopTypeController struct {
	svc *service.Service
}

// NewShopTypeController 创建店铺类型控制器实例
func NewShopTypeController(svc *service.Service) ShopTypeController {
	return &shopTypeController{svc: svc}
}

// 查询店铺类型列表
func (c *shopTypeController) GetShopTypeList(ctx *gin.Context) {
	res, err := c.svc.ShopTypeService.GetShopTypeList()
	if err != nil {
		result.Fail(ctx, constant.ErrCodeServerInternal, "查询店铺类型列表失败")
		return
	}
	result.Success(ctx, res)
}
