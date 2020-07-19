package model

import "github.com/jinzhu/gorm"

// 系统设置条目结构
type SysSettingItem struct {
	gorm.Model
	SysId uint   `json:"sys_id"`                            // 系统设置 ID
	Key   string `gorm:"varchar(100);not null;" json:"key"` // 键
	Val   string `gorm:"varchar(255);not null;" json:"val"` // 值
}
