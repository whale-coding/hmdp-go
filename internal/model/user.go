package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`        // 用户id
	Phone     string    `gorm:"column:phone;size:11;uniqueIndex" json:"phone"`       // 手机号码
	Password  string    `gorm:"column:password;size:128" json:"password"`            // 密码，加密存储
	NickName  string    `gorm:"column:nick_name;size:32" json:"nickName"`            // 昵称，默认是用户id
	Icon      string    `gorm:"column:icon;size:255" json:"icon"`                    // 人物头像
	CreatedAt time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 创建时间
	UpdatedAt time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"` // 更新时间
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (User) TableName() string {
	return "tb_user"
}

// UserDTO 用户数据传输对象
type UserDTO struct {
	ID       uint64 `json:"id"`       // 用户id
	NickName string `json:"nickName"` // 昵称
	Icon     string `json:"icon"`     // 人物头像
}

// LoginFormDTO 登录表单数据传输对象
type LoginFormDTO struct {
	Phone    string `json:"phone" binding:"required"` // 手机号，必填
	Code     string `json:"code" binding:"required"`  // 验证码，必填
	Password string `json:"password,omitempty"`       // 密码，可选字段，登录时不需要，注册时需要
}
