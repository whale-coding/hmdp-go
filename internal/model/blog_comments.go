package model

import (
	"time"
)

// BlogComments 博客评论表
type BlogComments struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`        // 主键
	UserID     uint64    `gorm:"column:user_id;not null" json:"userId"`               // 用户id
	BlogID     uint64    `gorm:"column:blog_id;not null" json:"blogId"`               // 探店id
	ParentID   uint64    `gorm:"column:parent_id;default:0;not null" json:"parentId"` // 关联的1级评论id，如果是一级评论，则值为0
	AnswerID   uint64    `gorm:"column:answer_id;default:0;not null" json:"answerId"` // 回复的评论id
	Content    string    `gorm:"column:content;size:255;not null" json:"content"`     // 回复的内容
	Liked      int       `gorm:"column:liked;default:0" json:"liked"`                 // 点赞数
	Status     uint8     `gorm:"column:status;default:0" json:"status"`               // 状态，0：正常，1：被举报，2：禁止查看
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"` // 更新时间
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (BlogComments) TableName() string {
	return "tb_blog_comments"
}
