package model

import "github.com/jinzhu/gorm"

// 页面结构
type Page struct {
	gorm.Model
	Title string `gorm:"varchar(100);not null;" json:"title"` // 标题
	Html  string `gorm:"MediumText;not null;" json:"html"`    // 页面内容
	Url   string `gorm:"varchar(100);not null;" json:"url"`   // 访问地址
}
