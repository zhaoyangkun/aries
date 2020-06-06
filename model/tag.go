package model

import (
	"aries/config/db"
	"github.com/jinzhu/gorm"
)

// 标签结构体
type Tag struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null;" json:"name"` // 标签名
	ArticleList []Article `gorm:"many2many:tag_article"`                   // 文章列表
}

// 初始化数据表
func init() {
	db.Db.AutoMigrate(&Tag{})
}
