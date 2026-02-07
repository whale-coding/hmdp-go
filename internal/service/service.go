package service

import (
	"hmdp-go/internal/repository"
)

// Service 服务层结构体
type Service struct {
	UserService           UserService
	UserInfoService       UserInfoService
	BlogService           BlogService
	BlogCommentsService   BlogCommentsService
	FollowService         FollowService
	ShopService           ShopService
	ShopTypeService       ShopTypeService
	VoucherService        VoucherService
	VoucherOrderService   VoucherOrderService
	SeckillVoucherService SeckillVoucherService
	SignService           SignService
}

// NewService 创建服务层
func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService:           NewUserService(repo),
		UserInfoService:       NewUserInfoService(repo),
		BlogService:           NewBlogService(repo),
		BlogCommentsService:   NewBlogCommentsService(repo),
		FollowService:         NewFollowService(repo),
		ShopService:           NewShopService(repo),
		ShopTypeService:       NewShopTypeService(repo),
		VoucherService:        NewVoucherService(repo),
		VoucherOrderService:   NewVoucherOrderService(repo),
		SeckillVoucherService: NewSeckillVoucherService(repo),
		SignService:           NewSignService(repo),
	}
}
