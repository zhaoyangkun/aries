package model

import (
	"aries/config/db"
	"aries/util"
	"github.com/jinzhu/gorm"
	"strings"
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
	var parentCategories []Category
	err := db.Db.Where("parent_category_id IS NULL").Find(&parentCategories).Error
	if err != nil {
		return categories, err
	}
	err = db.Db.Find(&categories).Error
	for i := range categories {
		if categories[i].ParentCategoryId > 0 {
			for j := range parentCategories {
				if parentCategories[j].ID == categories[i].ParentCategoryId {
					categories[i].ParentCategory = &parentCategories[j]
				}
			}
		}
	}
	return categories, err
}

// 获取分类数据（分页 +　搜索）
func (category Category) GetByPage(page *util.Pagination, key string) ([]Category, uint, error) {
	var parentCategories []Category
	var list []Category // 保存结果集
	// 获取父级分类
	err := db.Db.Where("parent_category_id IS NULL").Find(&parentCategories).Error
	if err != nil {
		return list, 0, err
	}
	// 创建语句
	query := db.Db.Model(&Category{}).Order("created_at desc", true)
	// 拼接搜索语句
	if key != "" {
		query = query.Where("name like concat('%',?,'%')", key)
	}
	// 分页
	total, err := util.ToPage(page, query, &list)
	// 拼接父子级分类
	for i := range list {
		if list[i].ParentCategoryId > 0 {
			for j := range parentCategories {
				if parentCategories[j].ID == list[i].ParentCategoryId {
					list[i].ParentCategory = &parentCategories[j]
				}
			}
		}
	}
	// 返回分页数据
	return list, total, err
}

// 添加分类
func (category Category) Create() error {
	err := db.Db.Create(&category).Error
	return err
}

// 修改分类
func (category Category) Update() error {
	return db.Db.Model(&category).Where("id = ?", category.ID).
		Updates(category).Error
}

// 删除分类
func (category Category) DeleteById() error {
	return db.Db.Where("id = ?", category.ID).
		Unscoped().Delete(&category).Error
}

// 批量删除分类
func (category Category) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",") // 根据 , 分割成字符串数组
	return db.Db.Where("id in (?)", idList).
		Unscoped().Delete(&category).Error
}
