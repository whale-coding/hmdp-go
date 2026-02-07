package model

import (
	"time"

	"gorm.io/gorm"
)

// SeckillVoucher 秒杀优惠券表
type SeckillVoucher struct {
	VoucherID  uint64         `gorm:"column:voucher_id;primaryKey" json:"voucherId"`       // 关联的优惠券的id，注意：非自增，需要程序生成
	Stock      int            `gorm:"column:stock;not null" json:"stock"`                  // 库存
	CreateTime time.Time      `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 创建时间
	BeginTime  time.Time      `gorm:"column:begin_time;not null" json:"beginTime"`         // 生效时间
	EndTime    time.Time      `gorm:"column:end_time;not null" json:"endTime"`             // 失效时间
	UpdateTime time.Time      `gorm:"column:update_time;autoUpdateTime" json:"updateTime"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`                                      // 软删除字段
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (SeckillVoucher) TableName() string {
	return "tb_seckill_voucher"
}
