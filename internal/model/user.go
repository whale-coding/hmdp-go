package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`        // 用户id
	Phone     string         `gorm:"column:phone;size:11;uniqueIndex" json:"phone"`       // 手机号码
	Password  string         `gorm:"column:password;size:128" json:"password"`            // 密码，加密存储
	NickName  string         `gorm:"column:nick_name;size:32" json:"nickName"`            // 昵称，默认是用户id
	Icon      string         `gorm:"column:icon;size:255" json:"icon"`                    // 人物头像
	CreatedAt time.Time      `gorm:"column:create_time;autoCreateTime" json:"createTime"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:update_time;autoUpdateTime" json:"updateTime"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                      // 软删除字段, gorm内置支持
}

func (User) TableName() string {
	return "tb_user"
}
