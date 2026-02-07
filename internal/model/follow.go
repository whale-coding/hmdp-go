package model

import "time"

// Follow 用户关注表
type Follow struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`        // 主键
	UserID       uint64    `gorm:"column:user_id;not null" json:"userId"`               // 用户id
	FollowUserID uint64    `gorm:"column:follow_user_id;not null" json:"followUserId"`  // 关联的用户id
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 创建时间
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (Follow) TableName() string {
	return "tb_follow"
}
