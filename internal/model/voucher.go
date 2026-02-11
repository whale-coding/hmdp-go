package model

import (
	"time"
)

// Voucher 代金券表
type Voucher struct {
	ID          uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`        // 主键
	ShopID      uint64     `gorm:"column:shop_id" json:"shopId"`                        // 商铺id
	Title       string     `gorm:"column:title;size:255;not null" json:"title"`         // 代金券标题
	SubTitle    string     `gorm:"column:sub_title;size:255" json:"subTitle"`           // 副标题
	Rules       string     `gorm:"column:rules;size:1024" json:"rules"`                 // 使用规则
	PayValue    uint64     `gorm:"column:pay_value;not null" json:"payValue"`           // 支付金额，单位是分
	ActualValue uint64     `gorm:"column:actual_value;not null" json:"actualValue"`     // 抵扣金额，单位是分
	Type        uint8      `gorm:"column:type;not null" json:"type"`                    // 0,普通券；1,秒杀券
	Status      uint8      `gorm:"column:status;not null" json:"status"`                // 1,上架; 2,下架; 3,过期
	Stock       *int       `gorm:"-" json:"stock,omitempty"`                            // 库存，非数据库字段
	BeginTime   *time.Time `gorm:"-" json:"beginTime,omitempty"`                        // 生效时间，非数据库字段
	EndTime     *time.Time `gorm:"-" json:"endTime,omitempty"`                          // 失效时间，非数据库字段
	CreateTime  time.Time  `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 创建时间
	UpdateTime  time.Time  `gorm:"column:update_time;autoUpdateTime" json:"updateTime"` // 更新时间
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (Voucher) TableName() string {
	return "tb_voucher"
}
