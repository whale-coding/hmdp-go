package model

import (
	"time"

	"gorm.io/gorm"
)

// Blog 探店博客表
type Blog struct {
	ID         uint64         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`        // 主键
	ShopID     uint64         `gorm:"column:shop_id;not null" json:"shopId"`               // 商户id
	UserID     uint64         `gorm:"column:user_id;not null" json:"userId"`               // 用户id
	Icon       string         `gorm:"-" json:"icon"`                                       // 用户图标，非数据库字段
	Name       string         `gorm:"-" json:"name"`                                       // 用户姓名，非数据库字段
	IsLike     bool           `gorm:"-" json:"isLike"`                                     // 是否点赞过了，非数据库字段
	Title      string         `gorm:"column:title;size:255;not null" json:"title"`         // 标题
	Images     string         `gorm:"column:images;size:2048;not null" json:"images"`      // 探店的照片，最多9张，多张以","隔开
	Content    string         `gorm:"column:content;size:2048;not null" json:"content"`    // 探店的文字描述
	Liked      int            `gorm:"column:liked;default:0" json:"liked"`                 // 点赞数量
	Comments   int            `gorm:"column:comments;default:0" json:"comments"`           // 评论数量
	CreateTime time.Time      `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 创建时间
	UpdateTime time.Time      `gorm:"column:update_time;autoUpdateTime" json:"updateTime"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`                                      // 软删除字段
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (Blog) TableName() string {
	return "tb_blog"
}
