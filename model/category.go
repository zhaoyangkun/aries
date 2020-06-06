package model

import (
	"aries/config/db"
	"github.com/jinzhu/gorm"
)

// 分类结构体
type Category struct {
	gorm.Model
	Type             uint      `gorm:"type:tinyint(1);unsigned;default:0" json:"type"`     // 分类类型，默认值为 0 表文章；1 表友链
	Name             string    `gorm:"type:varchar(100);not null;" json:"name"`            // 分类名称
	Url              string    `gorm:"type:varchar(100);not null;" json:"url"`             // 访问 URL
	ParentCategory   *Category `gorm:"foreignkey:ParentCategoryId" json:"parent_category"` // 父级分类
	ParentCategoryId uint      `json:"parent_category_id"`                                 // 父级分类 ID
}

// 初始化数据表
func init() {
	db.Db.AutoMigrate(&Category{})
}

// 获取所有分类
func (category Category) GetAll() []Category {
	var categories []Category
	db.Db.Find(&categories)
	return categories
}
