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
	idStr := ctx.Param("id")
	id, err := util.StringToUint64(idStr)
	if err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "商铺ID参数错误")
		return
	}

	res, err := c.svc.ShopService.GetShopById(id)
	if err != nil {
		log.Printf("查询商铺详情失败,商铺ID: %d,错误: %v\n", id, err)
		result.Fail(ctx, constant.ErrCodeServerInternal, "查询商铺详情失败")
		return
	}

	result.Success(ctx, res)
}
