package model

import (
	"aries/config/db"
	"aries/util"
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
func (category Category) GetAll() ([]Category, error) {
	var categories []Category
	err := db.Db.Select("`categories`.*,`ParentCategory`.*").
		Joins("left join `categories` as `ParentCategory` " +
			"on `categories`.`parent_category_id` = `ParentCategory`.`id`").
		Find(&categories).Error
	return categories, err
}

// 获取分类数据（分页 +　搜索）
func (category Category) GetByPage(page *util.Pagination, key string) ([]Category, uint, error) {
	var list []Category // 保存结果集
	// 创建语句
	query := db.Db.Model(&Category{}).Preload("ParentCategory").
		Order("created_at desc", true)
	// 拼接搜索语句
	if key != "" {
		query = query.Where("name like concat('%',?,'%')", key)
	}
	// 分页
	total, err := util.ToPage(page, query, &list)
	// 返回分页数据
	return list, total, err
}
