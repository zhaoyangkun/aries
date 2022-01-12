package models

import (
	"aries/config/db"
	"aries/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

// Gallery 图库
type Gallery struct {
	gorm.Model
	Category   Category `gorm:"ForeignKey:CategoryId"`                   // 分类
	CategoryId uint     `json:"category_id"`                             // 分类 ID
	URL        string   `gorm:"type:varchar(255);not null;" json:"url"`  // 图片链接
	Name       string   `gorm:"type:varchar(255);not null;" json:"name"` // 图片名称
	Desc       string   `gorm:"type:varchar(255);" json:"desc"`          // 图片描述
	Location   string   `gorm:"type:varchar(50);" json:"location"`       // 拍摄地点
}

// GetAll 获取所有图库
func (Gallery) GetAll() (list []Gallery, err error) {
	err = db.Db.Preload("Category").Order("created_at desc", true).Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetById 根据 ID 获取图库
func (Gallery) GetById(id uint) (gallery Gallery, err error) {
	err = db.Db.Where("`id` = ?", id).First(&gallery).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByPage 分页获取图库
func (g Gallery) GetByPage(page *utils.Pagination, categoryId uint, key string) (list []Gallery, total uint, err error) {
	query := db.Db.Preload("Category").Order("created_at desc", true).Find(&list)

	if categoryId > 0 {
		query = query.Where("`category_id` = ?", categoryId)
	}

	if key != "" {
		query = query.Where("`name` like concat('%',?,'%') or `desc` like concat('%',?,'%')", key, key)
	}

	total, err = utils.ToPage(page, query, &list)
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// Create 创建图库
func (g *Gallery) Create() error {
	return db.Db.Create(&g).Error
}

// Update 更新图库
func (g *Gallery) Update() error {
	return db.Db.Model(&Gallery{}).Updates(&g).Error
}

// MultiDelByIds 批量删除图库
func (Gallery) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",")

	return db.Db.Where("`id` in (?)", idList).Unscoped().Delete(&Gallery{}).Error
}
