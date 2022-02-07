package models

import (
	"aries/config/db"
	"aries/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

// Tag 标签
type Tag struct {
	gorm.Model
	ArticleList []Article `gorm:"many2many:tag_article" json:"article_list"` // 文章列表
	Name        string    `gorm:"type:varchar(100);not null;" json:"name"`   // 标签名
	Count       uint      `gorm:"type:int;default:0;" json:"count"`          // 文章数量
}

// GetAll 获取所有标签
func (Tag) GetAll() ([]Tag, error) {
	var list []Tag
	err := db.Db.Preload("ArticleList").Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return list, err
}

// GetAllWithNoArticle 获取所有标签，不包含文章列表
func (Tag) GetAllWithNoArticle() ([]Tag, error) {
	var list []Tag
	err := db.Db.Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return list, err
}

// GetById 根据主键获取标签
func (Tag) GetById(id string) (Tag, error) {
	var t Tag
	err := db.Db.Where("`id` = ?", id).First(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return t, err
}

// GetByName 根据名称获取标签
func (Tag) GetByName(name string) (tag Tag, err error) {
	err = db.Db.Where("`name` = ?", name).First(&tag).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByPage 分页获取标签
func (tag Tag) GetByPage(page *utils.Pagination, key string) ([]Tag, uint, error) {
	var list []Tag

	query := db.Db.Model(&Tag{}).Order("created_at desc", true)

	if key != "" {
		query = query.Where("`name` like concat('%',?,'%')", key)
	}

	total, err := utils.ToPage(page, query, &list)
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return list, total, err
}

// Create 添加标签
func (tag *Tag) Create() error {
	return db.Db.Create(&tag).Error
}

// Update 更新标签
func (tag *Tag) Update() error {
	return db.Db.Model(&Tag{}).Updates(tag).Error
}

// DeleteById 删除标签
func (Tag) DeleteById(id string) error {
	// 开启事务
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	// 删除标签文章表中的记录
	err := tx.Exec("delete from `tag_article` where `tag_id` = ?", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除标签表中的记录
	err = tx.Where("`id` = ?", id).Unscoped().Delete(&Tag{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// MultiDelByIds 批量删除标签
func (Tag) MultiDelByIds(ids string) error {
	// 开始事务
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
	// 删除标签文章表中的记录
	err := tx.Exec("delete from `tag_article` where `tag_id` in (?)", idList).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除标签表中的记录
	err = tx.Where("`id` in (?)", idList).Unscoped().Delete(&Tag{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
