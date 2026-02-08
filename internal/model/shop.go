package model

import (
	"time"
)

// Shop 商铺表
type Shop struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`        // 主键
	Name       string    `gorm:"column:name;size:128;not null" json:"name"`           // 商铺名称
	TypeID     uint64    `gorm:"column:type_id;not null" json:"typeId"`               // 商铺类型的id
	Images     string    `gorm:"column:images;size:1024;not null" json:"images"`      // 商铺图片，多个图片以','隔开
	Area       string    `gorm:"column:area;size:128" json:"area"`                    // 商圈，例如陆家嘴
	Address    string    `gorm:"column:address;size:255;not null" json:"address"`     // 地址
	X          float64   `gorm:"column:x;not null" json:"x"`                          // 经度
	Y          float64   `gorm:"column:y;not null" json:"y"`                          // 维度
	AvgPrice   uint64    `gorm:"column:avg_price" json:"avgPrice"`                    // 均价，取整数
	Sold       int       `gorm:"column:sold;default:0;not null" json:"sold"`          // 销量
	Comments   int       `gorm:"column:comments;default:0;not null" json:"comments"`  // 评论数量
	Score      int       `gorm:"column:score;default:0;not null" json:"score"`        // 评分，1~5分，乘10保存，避免小数
	OpenHours  string    `gorm:"column:open_hours;size:32" json:"openHours"`          // 营业时间，例如 10:00-22:00
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"` // 更新时间
	Distance   float64   `gorm:"-" json:"distance"`                                   // 距离，非数据库字段
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (Shop) TableName() string {
	return "tb_shop"
}
