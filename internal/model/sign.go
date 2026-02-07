package model

// Sign 签到记录表
type Sign struct {
	ID       uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 主键
	UserID   uint64 `gorm:"column:user_id;not null" json:"userId"`        // 用户id
	Year     uint16 `gorm:"column:year;not null" json:"year"`             // 签到的年
	Month    uint8  `gorm:"column:month;not null" json:"month"`           // 签到的月
	Date     string `gorm:"column:date;type:date;not null" json:"date"`   // 签到的日期
	IsBackup uint8  `gorm:"column:is_backup;default:0" json:"isBackup"`   // 是否补签，0：否，1：是
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (Sign) TableName() string {
	return "tb_sign"
}
