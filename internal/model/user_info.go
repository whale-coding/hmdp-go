package model

import (
	"time"

	"gorm.io/gorm"
)

// UserInfo 用户信息表
type UserInfo struct {
	UserID     uint64         `gorm:"column:user_id;primaryKey;autoIncrement" json:"userId"` // 用户id
	City       string         `gorm:"column:city;type:varchar(64)" json:"city"`              // 城市
	Introduce  string         `gorm:"column:introduce;type:varchar(128)" json:"introduce"`   // 个人介绍
	Fans       int            `gorm:"column:fans;default:0" json:"fans"`                     // 粉丝数量
	Followee   int            `gorm:"column:followee;default:0" json:"followee"`             // 关注的人的数量
	Gender     uint8          `gorm:"column:gender;default:0" json:"gender"`                 // 性别，0：男，1：女
	Birthday   *time.Time     `gorm:"column:birthday;type:date" json:"birthday"`             // 生日
	Credits    int            `gorm:"column:credits;default:0" json:"credits"`               // 积分
	Level      uint8          `gorm:"column:level;default:0" json:"level"`                   // 会员级别，0~9级，0代表未开通会员
	CreateTime time.Time      `gorm:"column:create_time;autoCreateTime" json:"createTime"`   // 创建时间
	UpdateTime time.Time      `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`   // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`                                        // 软删除字段, gorm内置支持
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (UserInfo) TableName() string {
	return "tb_user_info"
}
