package repository

import (
	"hmdp-go/internal/model"
	"time"

	"gorm.io/gorm"
)

// VoucherRepository 代金券仓库接口
type VoucherRepository interface {
	FindByShopID(shopID uint64) ([]model.Voucher, error)
}

// voucherRepository 代金券仓库实现
type voucherRepository struct {
	db *gorm.DB
}

// NewVoucherRepository 创建代金券仓库实例
func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db: db}
}

// VoucherWithSeckill 用于接收联表查询结果
type VoucherWithSeckill struct {
	model.Voucher
	Stock     *int       `gorm:"column:stock"`
	BeginTime *time.Time `gorm:"column:begin_time"`
	EndTime   *time.Time `gorm:"column:end_time"`
}

// FindByShopID 根据店铺ID查询代金券（LEFT JOIN 秒杀券表）
func (r *voucherRepository) FindByShopID(shopID uint64) ([]model.Voucher, error) {
	var results []VoucherWithSeckill

	err := r.db.Table("tb_voucher v").
		Select(`v.id, v.shop_id, v.title, v.sub_title, v.rules, v.pay_value, 
				v.actual_value, v.type, v.status, v.create_time, v.update_time,
				sv.stock, sv.begin_time, sv.end_time`).
		Joins("LEFT JOIN tb_seckill_voucher sv ON v.id = sv.voucher_id").
		Where("v.shop_id = ? AND v.status = 1", shopID).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// 转换为 Voucher 切片，并填充秒杀券字段
	vouchers := make([]model.Voucher, len(results))
	for i, r := range results {
		vouchers[i] = r.Voucher
		// 只有秒杀券才填充额外字段
		if r.Type == 1 {
			vouchers[i].Stock = r.Stock
			vouchers[i].BeginTime = r.BeginTime
			vouchers[i].EndTime = r.EndTime
		}
	}

	return vouchers, nil
}
