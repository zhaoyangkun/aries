package models

import (
	"aries/config/db"
	"aries/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

// 日志
type Journal struct {
	gorm.Model
	IsSecret bool   `gorm:"type:bool;default:false;" json:"is_secret"` // 是否私密
	Content  string `gorm:"type:varchar(255);not null;" json:"content"`     // 内容
}

// 获取所有日志
func (Journal) GetAll() (list []Journal, err error) {
	err = db.Db.Order("created_at desc", true).Find(&list).Error

	return
}

// 根据 ID 获取日志
func (Journal) GetById(id uint) (journal Journal, err error) {
	err = db.Db.Where("`id` = ?", id).First(&journal).Error

	return
}

// 分页获取日志
func (Journal) GetByPage(page *utils.Pagination, key string) (list []Journal, total uint, err error) {
	query := db.Db.Order("created_at desc", true).Find(&list)

	if key != "" {
		query = query.Where("`content` like concat('%',?,'%')", key)
	}

	total, err = utils.ToPage(page, query, &list)

	return
}

// 创建日志
func (j *Journal) Create() error {
	return db.Db.Create(&j).Error
}

// 更新日志
func (j *Journal) Update() error {
	return db.Db.Model(&Journal{}).Where("`id` = ?", j.ID).Updates(map[string]interface{}{
		"is_secret": j.IsSecret,
		"content":   j.Content,
	}).Error
}

// 批量删除日志
func (Journal) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",")

	return db.Db.Where("`id` in (?)", idList).Unscoped().Delete(&Journal{}).Error
}
