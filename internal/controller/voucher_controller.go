package controller

import (
	"hmdp-go/internal/constant"
	"hmdp-go/internal/service"
	"hmdp-go/pkg/result"
	"hmdp-go/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

// VoucherController 代金券控制器接口
type VoucherController interface {
	AddCommonVoucher(ctx *gin.Context)
	AddSeckillVoucher(ctx *gin.Context)
	GetVoucherByShopId(ctx *gin.Context)
}

// voucherController 代金券控制器实现
type voucherController struct {
	svc *service.Service
}

// NewVoucherController 创建代金券控制器实例
func NewVoucherController(svc *service.Service) VoucherController {
	return &voucherController{svc: svc}
}

// 新增普通券
func (c *voucherController) AddCommonVoucher(ctx *gin.Context) {
	// 获取参数
	// var voucherDTO model.VoucherDTO
	// if err := ctx.ShouldBind(&voucherDTO); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, model.Result{
	// 		Code: model.CodeInvalidParam,
	// 		Msg:  "参数错误",
	// 	})
	// 	return
	// }
}

// 新增秒杀券
func (c *voucherController) AddSeckillVoucher(ctx *gin.Context) {
	// 获取参数
	// var voucherDTO model.VoucherDTO
	// if err := ctx.ShouldBind(&voucherDTO); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, model.Result{
	// 		Code: model.CodeInvalidParam,
	// 		Msg:  "参数错误",
	// 	})
	// 	return
	// }
}

// 查询店铺的优惠券列表
func (c *voucherController) GetVoucherByShopId(ctx *gin.Context) {
	// 获取店铺ID参数（路径参数）
	shopIdStr := ctx.Param("id")
	log.Println("shopIdStr:", shopIdStr)
	shopId, err := util.StringToUint64(shopIdStr)
	if err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "店铺ID参数错误")
		return
	}
	log.Printf("查询店铺优惠券列表,店铺ID: %d\n", shopId)

	// 调用服务层查询优惠券列表
	vouchers, err := c.svc.VoucherService.GetVoucherByShopId(shopId)
	if err != nil {
		log.Printf("查询店铺优惠券列表失败,店铺ID: %d,错误: %v\n", shopId, err)
		result.Fail(ctx, constant.ErrCodeServerInternal, "查询优惠券列表失败")
		return
	}

	result.Success(ctx, vouchers)
}
