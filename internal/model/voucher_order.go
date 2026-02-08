package model

import (
	"time"
)

// VoucherOrder 代金券订单表
type VoucherOrder struct {
	ID         uint64     `gorm:"column:id;primaryKey" json:"id"`                      // 主键，注意：非自增，需要程序生成
	UserID     uint64     `gorm:"column:user_id;not null" json:"userId"`               // 下单的用户 id
	VoucherID  uint64     `gorm:"column:voucher_id;not null" json:"voucherId"`         // 购买的代金券 id
	PayType    uint8      `gorm:"column:pay_type;not null" json:"payType"`             // 支付方式 1：余额支付；2：支付宝；3：微信
	Status     uint8      `gorm:"column:status;not null" json:"status"`                // 订单状态，1：未支付；2：已支付；3：已核销；4：已取消；5：退款中；6：已退款
	CreateTime time.Time  `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 下单时间
	PayTime    *time.Time `gorm:"column:pay_time" json:"payTime"`                      // 支付时间
	UseTime    *time.Time `gorm:"column:use_time" json:"useTime"`                      // 核销时间
	RefundTime *time.Time `gorm:"column:refund_time" json:"refundTime"`                // 退款时间
	UpdateTime time.Time  `gorm:"column:update_time;autoUpdateTime" json:"updateTime"` // 更新时间
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (VoucherOrder) TableName() string {
	return "tb_voucher_order"
}
