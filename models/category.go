package models

import (
	"aries/config/db"
	"aries/utils"

	"strings"

	"github.com/jinzhu/gorm"
)

// Category 分类
type Category struct {
	gorm.Model
	Children []*Category `gorm:"ForeignKey:ParentId" json:"children"`            // 子级分类列表
	ParentId uint        `json:"parent_id"`                                      // 父类 ID：0 表父级分类，大于 0 表子级分类
	Type     uint        `gorm:"type:tinyint(1);unsigned;default:0" json:"type"` // 分类类型，默认值为 0 表文章；1 表友链,2 表示图库
	Name     string      `gorm:"type:varchar(100);not null;" json:"name"`        // 分类名称
	Url      string      `gorm:"type:varchar(100)" json:"url"`                   // 访问 URL
	Count    uint        `gorm:"type:int;default:0;" json:"count"`               // 文章数量
}

// GetAllByType 根据类别获取所有分类
func (category Category) GetAllByType(categoryType uint) ([]Category, error) {
	var categories []Category
	var children []Category

	// 查询子分类
	err := db.Db.Where("`parent_id` > 0").Find(&children).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return categories, err
	}

	// 根据类别查询分类，若 categoryType > 2，表示查询所有分类
	if categoryType <= 2 {
		err = db.Db.Where("`type` = ?", categoryType).Find(&categories).Error
	} else {
		err = db.Db.Find(&categories).Error
	}
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return categories, err
	}

	// 将子分类并入父分类
	for i := range categories {
		if categories[i].ParentId == 0 {
			for j := range children {
				if children[j].ParentId == categories[i].ID {
					categories[i].Children = append(categories[i].Children, &children[j])
				}
			}
		}
	}

	return categories, err
}

// GetGalleryCategories 获取有图库的分类
func (category Category) GetGalleryCategories() (list []Category, err error) {
	err = db.Db.Where("`id` in (select `category_id` from `galleries` group by `category_id`)").Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByPage 获取分类数据（分页 +　搜索）
func (category Category) GetByPage(page *utils.Pagination, key string, categoryType uint) ([]Category, uint, error) {
	var list []Category // 保存结果集
	var children []Category

	// 查询子分类
	err := db.Db.Where("`parent_id` > 0").Find(&children).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return list, 0, err
	}

	// 创建语句
	query := db.Db.Model(&Category{}).Where("`type` = ?", categoryType).
		Order("created_at desc", true)

	// 拼接搜索语句
	if key != "" {
		query = query.Where("`name` like concat('%',?,'%')", key)
	}

	// 分页
	total, err := utils.ToPage(page, query, &list)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return list, 0, err
	}

	// 将子类并入父类
	for i := range list {
		if list[i].ParentId == 0 {
			for j := range children {
				if children[j].ParentId == list[i].ID {
					list[i].Children = append(list[i].Children, &children[j])
				}
			}
		}
	}

	// 返回分页数据
	return list, total, err
}

// GetAllParents 获取所有父类
func (category Category) GetAllParents(categoryType uint) (list []Category, err error) {
	err = db.Db.Where("`parent_id` = 0 and `type` = ?", categoryType).Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByName 根据分类名称获取分类
func (Category) GetByName(name string, categoryType uint) (category Category, err error) {
	err = db.Db.Where("`name` = ? and `type` = ?", name, categoryType).First(&category).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByUrl 根据 URL 获取分类
func (Category) GetByUrl(url string) (category Category, err error) {
	err = db.Db.Where("`url` = ?", url).First(&category).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// Create 添加分类
func (category *Category) Create() error {
	return db.Db.Create(&category).Error
}

// Update 修改分类
func (category *Category) Update() error {
	// 使用 map 来更新，避免 gorm 默认不更新值为 nil, false, 0 的字段
	return db.Db.Model(&category).
		Updates(map[string]interface{}{
			"name":      category.Name,
			"parent_id": category.ParentId,
			"type":      category.Type,
			"url":       category.Url,
		}).Error
}

// DeleteById 删除分类
func (category Category) DeleteById(id uint) error {
	tx := db.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	//  更新关联文章
	err := tx.Model(&Article{}).Where("`category_id` = ?", id).Update("category_id", nil).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Exec("update `categories` set `parent_id` = 0 where `parent_id` = ?", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("`id` = ?", id).Unscoped().Delete(&category).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// MultiDelByIds 批量删除分类
func (category Category) MultiDelByIds(ids string) error {
	tx := db.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	idList := strings.Split(ids, ",") // 根据 , 分割成字符串数组
	// 更新关联文章
	err := tx.Model(&Article{}).Where("`category_id` in (?)", idList).
		Update("category_id", nil).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Exec("update `categories` set `parent_id` = 0 where `parent_id` in (?)", idList).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("`id` in (?)", idList).Unscoped().Delete(&category).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
