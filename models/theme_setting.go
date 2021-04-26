package models

import "github.com/jinzhu/gorm"

// ThemeSetting 主题设置
type ThemeSetting struct {
	gorm.Model
	ThemeID string `gorm:"type:varchar(50);not null;" json:"theme_id"` // 主题 ID
	Key     string `gorm:"type:varchar(100);not null;" json:"key"`     // 键
	Val     string `gorm:"type:varchar(255);not null;" json:"val"`     // 值
}
