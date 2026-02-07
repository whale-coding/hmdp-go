package controller

import (
	"hmdp-go/internal/service"
)

// Controller 控制器结构体
type Controller struct {
	UserController           UserController
	UserInfoController       UserInfoController
	BlogController           BlogController
	BlogCommentsController   BlogCommentsController
	FollowController         FollowController
	ShopController           ShopController
	ShopTypeController       ShopTypeController
	VoucherController        VoucherController
	VoucherOrderController   VoucherOrderController
	SeckillVoucherController SeckillVoucherController
	SignController           SignController
}

// NewController 创建控制器
func NewController(svc *service.Service) *Controller {
	return &Controller{
		UserController:           NewUserController(svc),
		UserInfoController:       NewUserInfoController(svc),
		BlogController:           NewBlogController(svc),
		BlogCommentsController:   NewBlogCommentsController(svc),
		FollowController:         NewFollowController(svc),
		ShopController:           NewShopController(svc),
		ShopTypeController:       NewShopTypeController(svc),
		VoucherController:        NewVoucherController(svc),
		VoucherOrderController:   NewVoucherOrderController(svc),
		SeckillVoucherController: NewSeckillVoucherController(svc),
		SignController:           NewSignController(svc),
	}
}
