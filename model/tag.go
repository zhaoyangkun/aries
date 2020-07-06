package model

import (
	"aries/config/db"
	"aries/util"
	"github.com/jinzhu/gorm"
	"strings"
)

// 标签结构
type Tag struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null;" json:"name"` // 标签名
	ArticleList []Article `gorm:"many2many:tag_article"`                   // 文章列表
}

// 获取所有标签
func (Tag) GetAll() ([]Tag, error) {
	var list []Tag
	err := db.Db.Preloads("ArticleList").Find(&list).Error
	return list, err
}

// 根据主键获取标签
func (Tag) GetById(id string) (Tag, error) {
	var t Tag
	err := db.Db.Where("id = ?", id).First(&t).Error
	return t, err
}

// 分页获取标签
func (tag Tag) GetByPage(page *util.Pagination, key string) ([]Tag, uint, error) {
	var list []Tag
	query := db.Db.Preloads("ArticleList").Model(&Tag{}).
		Order("created_at desc", true)
	if key != "" {
		query = query.Where("name like concat('%',?,'%')", key)
	}
	total, err := util.ToPage(page, query, &list)
	return list, total, err
}

// 添加标签
func (tag Tag) Create() error {
	return db.Db.Create(&tag).Error
}

// 更新标签
func (tag Tag) Update() error {
	return db.Db.Model(&Tag{}).Updates(&tag).Error
}

// 删除标签
func (Tag) DeleteById(id string) error {
	// 删除标签文章表中的记录
	err := db.Db.Exec("delete from tag_article where tag_id = ?", id).Error
	if err != nil {
		return err
	}
	// 删除标签表中的记录
	return db.Db.Where("id = ?", id).Unscoped().Delete(&Tag{}).Error
}

// 批量删除标签
func (Tag) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",") // 根据 , 分割成字符串数组
	// 删除标签文章表中的记录
	err := db.Db.Exec("delete from tag_article where tag_id in (?)", idList).Error
	if err != nil {
		return err
	}
	// 删除标签表中的记录
	return db.Db.Where("id in (?)", idList).Unscoped().Delete(&Tag{}).Error
}
