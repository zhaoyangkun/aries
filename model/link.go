package model

import (
	"github.com/jinzhu/gorm"
)

// 友情链接结构
type Link struct {
	gorm.Model
	Category   Category `gorm:"ForeignKey:CategoryId" json:"category"` // 分类
	CategoryId uint     `json:"category_id"`                           // 分类 ID
	Name       string   `gorm:"varchar(100);not null;" json:"name"`    // 网站名称
	Url        string   `gorm:"varchar(255);not null;" json:"url"`     // 网站地址
	Desc       string   `gorm:"varchar(255);" json:"desc"`             // 网站描述
	Icon       string   `gorm:"varchar(255);not null;" json:"icon"`    // 图标
}
