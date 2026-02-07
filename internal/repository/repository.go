package repository

import (
	"gorm.io/gorm"
)

// Repository 仓库结构体,仓库统一聚合体，业务层通过它访问各个具体的仓库
type Repository struct {
	UserRepo           UserRepository
	UserInfoRepo       UserInfoRepository
	BlogRepo           BlogRepository
	BlogCommentsRepo   BlogCommentsRepository
	FollowRepo         FollowRepository
	ShopRepo           ShopRepository
	ShopTypeRepo       ShopTypeRepository
	VoucherRepo        VoucherRepository
	VoucherOrderRepo   VoucherOrderRepository
	SeckillVoucherRepo SeckillVoucherRepository
	SignRepo           SignRepository
}

// NewRepository 构造函数，注入所有具体的仓库实现
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepo:           NewUserUserRepository(db),
		UserInfoRepo:       NewUserInfoRepository(db),
		BlogRepo:           NewBlogRepository(db),
		BlogCommentsRepo:   NewBlogCommentsRepository(db),
		FollowRepo:         NewFollowRepository(db),
		ShopRepo:           NewShopRepository(db),
		ShopTypeRepo:       NewShopTypeRepository(db),
		VoucherRepo:        NewVoucherRepository(db),
		VoucherOrderRepo:   NewVoucherOrderRepository(db),
		SeckillVoucherRepo: NewSeckillVoucherRepository(db),
		SignRepo:           NewSignRepository(db),
	}
}
