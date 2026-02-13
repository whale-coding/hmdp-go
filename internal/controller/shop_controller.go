package controller

import (
	"hmdp-go/internal/constant"
	"hmdp-go/internal/model"
	"hmdp-go/internal/service"
	"hmdp-go/pkg/result"
	"hmdp-go/pkg/util"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShopController 商铺控制器接口
type ShopController interface {
	QueryShopByType(ctx *gin.Context)
	GetShopById(ctx *gin.Context)
	CreateShop(ctx *gin.Context)
	UpdateShop(ctx *gin.Context)
}

// shopController 商铺控制器实现
type shopController struct {
	svc *service.Service
}

// NewShopController 创建商铺控制器实例
func NewShopController(svc *service.Service) ShopController {
	return &shopController{svc: svc}
}

// QueryShopByType 根据商铺类型分页查询商铺信息
func (c *shopController) QueryShopByType(ctx *gin.Context) {
	// 绑定查询参数
	typeIdStr := ctx.Query("typeId")
	currentStr := ctx.DefaultQuery("current", "1")
	xStr := ctx.Query("x")
	yStr := ctx.Query("y")

	// 参数校验 - typeId 必填
	typeId, err := util.StringToUint64(typeIdStr)
	if err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "商铺类型ID参数错误")
		return
	}

	// x, y 可选（用于计算距离）
	var x, y float64
	if xStr != "" && yStr != "" {
		x, _ = strconv.ParseFloat(xStr, 64)
		y, _ = strconv.ParseFloat(yStr, 64)
	}

	// 解析分页参数 current
	var page model.PaginationRequest
	if current, err := strconv.Atoi(currentStr); err == nil && current > 0 {
		page.PageNo = current
	}
	page.SetDefault()

	log.Printf("查询商铺列表,类型ID: %d,页码: %d,坐标: (%f, %f)\n", typeId, page.PageNo, x, y)

	// 调用服务层查询商铺列表
	res, err := c.svc.ShopService.QueryShopByType(typeId, &page, x, y)
	if err != nil {
		log.Printf("查询商铺列表失败,商铺类型ID: %d,错误: %v\n", typeId, err)
		result.Fail(ctx, constant.ErrCodeServerInternal, "查询商铺列表失败")
		return
	}

	result.Success(ctx, res)
}

// GetShopById 根据id查询商铺信息
func (c *shopController) GetShopById(ctx *gin.Context) {
	// 从路径参数获取商铺ID
	idStr := ctx.Param("id")
	id, err := util.StringToUint64(idStr)
	if err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "商铺ID参数错误")
		return
	}
	// 调用服务层查询商铺详情
	res, err := c.svc.ShopService.GetShopById(id)
	if err != nil {
		log.Printf("查询商铺详情失败,商铺ID: %d,错误: %v\n", id, err)
		result.Fail(ctx, constant.ErrCodeServerInternal, "查询商铺详情失败")
		return
	}

	result.Success(ctx, res)
}

// 新增商铺信息
func (c *shopController) CreateShop(ctx *gin.Context) {
	// 获取参数
	var shop model.Shop
	if err := ctx.ShouldBindJSON(&shop); err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "请求参数错误")
		return
	}

	// 调用服务层创建商铺
	err := c.svc.ShopService.CreateShop(&shop)
	if err != nil {
		log.Printf("创建商铺失败,商铺信息: %+v,错误: %v\n", shop, err)
		result.Fail(ctx, constant.ErrCodeServerInternal, "创建商铺失败")
		return
	}

	result.Success(ctx, nil)
}

// 修改商铺信息
func (c *shopController) UpdateShop(ctx *gin.Context) {
	// 获取参数
	var shop model.Shop
	if err := ctx.ShouldBindJSON(&shop); err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "请求参数错误")
		return
	}

	// 调用服务层更新商铺
	err := c.svc.ShopService.UpdateShop(&shop)
	if err != nil {
		log.Printf("更新商铺失败,商铺信息: %+v,错误: %v\n", shop, err)
		result.Fail(ctx, constant.ErrCodeServerInternal, "更新商铺失败")
		return
	}

	result.Success(ctx, nil)
}
