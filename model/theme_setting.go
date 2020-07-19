package model

import "github.com/jinzhu/gorm"

type ThemeSetting struct {
	gorm.Model
	ThemeID uint   `json:"theme_id"`                          // 主题 ID
	Key     string `gorm:"varchar(100);not null;" json:"key"` // 键
	Val     string `gorm:"varchar(255);not null;" json:"val"` // 值
}
