package models

import (
	"aries/config/db"
	"aries/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

// Link 友情链接
type Link struct {
	gorm.Model
	Category   Category `gorm:"ForeignKey:CategoryId" json:"category"`   // 分类
	CategoryId uint     `json:"category_id"`                             // 分类 ID
	Name       string   `gorm:"type:varchar(100);not null;" json:"name"` // 网站名称
	Url        string   `gorm:"type:varchar(255);not null;" json:"url"`  // 网站地址
	Desc       string   `gorm:"type:varchar(255);" json:"desc"`          // 网站描述
	Icon       string   `gorm:"type:varchar(255);not null;" json:"icon"` // 图标
}

// GetAll 获取所有友链
func (Link) GetAll() (list []Link, err error) {
	err = db.Db.Preload("Category").Order("category_id desc,created_at desc", true).Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByPage 分页获取友链
func (Link) GetByPage(page *utils.Pagination, key string, categoryId uint) ([]Link, uint, error) {
	var list []Link

	query := db.Db.Model(&Link{}).Preload("Category").Order("category_id desc,created_at desc", true)

	if key != "" {
		query = query.Where("`name` like concat('%',?,'%')", key)
	}

	if categoryId > 0 {
		query = query.Where("`category_id` = ?", categoryId)
	}

	total, err := utils.ToPage(page, query, &list)
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return list, total, err
}

// GetById 根据 ID 获取友链
func (Link) GetById(id string) (link Link, err error) {
	err = db.Db.Preload("Category").Where("`id` = ?", id).First(&link).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// Create 添加友链
func (link *Link) Create() (err error) {
	err = db.Db.Create(&link).Error
	return
}

// Update 更新友链
func (link *Link) Update() (err error) {
	err = db.Db.Model(&Link{}).Where("`id` = ?", link.ID).
		Updates(map[string]interface{}{
			"category_id": link.CategoryId,
			"name":        link.Name,
			"url":         link.Url,
			"desc":        link.Desc,
			"icon":        link.Icon,
		}).Error

	return
}

// DeleteById 删除友链
func (Link) DeleteById(id string) error {
	return db.Db.Where("`id` = ?", id).Unscoped().Delete(&Link{}).Error
}

// MultiDelByIds 批量删除友链
func (Link) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",")
	return db.Db.Where("`id` in (?)", idList).Unscoped().Delete(&Link{}).Error
}
