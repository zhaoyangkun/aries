package models

import (
	"github.com/jinzhu/gorm"
)

// 主题
type Theme struct {
	gorm.Model
	ThemeInfo
	IsUsed bool `gorm:"type:bool;default:false;" json:"is_used"`
}

// 主题信息
type ThemeInfo struct {
	AuthorInfo `yaml:"author"`
	ThemeID    string `gorm:"type:varchar(30);not null;" yaml:"id" json:"theme_id"`
	ThemeName  string `gorm:"type:varchar(100);not null;" yaml:"name" json:"theme_name"`
	Desc       string `gorm:"type:varchar(255);" yaml:"desc" json:"desc"`
	Image      string `gorm:"type:varchar(255);not null;" yaml:"image" json:"image"`
	Repo       string `gorm:"type:varchar(255);" yaml:"repo" json:"repo"`
	Version    string `gorm:"type:varchar(30);not null;" yaml:"version" json:"version"`
}

// 作者信息
type AuthorInfo struct {
	AuthorName string `gorm:"type:varchar(30);not null;" yaml:"name" json:"author_name"`
	Website    string `gorm:"type:varchar(255);" yaml:"website" json:"website"`
	HeadImg    string `gorm:"type:varchar(255);" yaml:"head_img" json:"head_img"`
}
