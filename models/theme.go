package models

import (
	"github.com/jinzhu/gorm"
)

// 主题
type Theme struct {
	gorm.Model
	ThemeName string `gorm:"type:varchar(100);not null;" json:"theme_name"` // 主题名称
	Author    string `gorm:"type:varchar(30);not null;" json:"author"`      // 作者
	IsUsed    bool   `gorm:"type:bool;default:false;" json:"is_used"`       // 是否启用，默认 false
}
