package model

import (
	"time"
)

// ShopType 店铺类型表
type ShopType struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 主键
	Name       string    `gorm:"column:name;size:32" json:"name"`              // 类型名称
	Icon       string    `gorm:"column:icon;size:255" json:"icon"`             // 图标
	Sort       int       `gorm:"column:sort" json:"sort"`                      // 顺序
	CreateTime time.Time `gorm:"column:create_time" json:"-"`                  // 创建时间，不序列化
	UpdateTime time.Time `gorm:"column:update_time" json:"-"`                  // 更新时间，不序列化
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (ShopType) TableName() string {
	return "tb_shop_type"
}
