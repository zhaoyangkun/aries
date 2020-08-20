package models

import (
	"aries/config/db"
	"aries/utils"

	"github.com/jinzhu/gorm"
)

// 图片
type Picture struct {
	gorm.Model
	StorageType string `gorm:"varchar(20);not null" json:"storage_type"` // 存储类型
	Hash        string `gorm:"varchar(100)" json:"hash"`                 // 标识
	FileName    string `gorm:"varchar(255);not null" json:"file_name"`   // 文件名
	URL         string `gorm:"varchar(255);not null" json:"url"`         // 访问地址
	Size        uint   `gorm:"int(6)" json:"size"`                       // 空间大小（KB）
}

// 分页获取图片
func (Picture) GetByPage(page *utils.Pagination, key string, storageType string) (list []Picture, total uint, err error) {
	query := db.Db.Model(&Picture{}).Order("created_at desc", true)

	if key != "" {
		query = query.Where("`file_name` like concat('%',?,'%')", key)
	}

	if storageType != "" {
		query = query.Where("`storage_type` = ?", storageType)
	}

	// 分页
	total, err = utils.ToPage(page, query, &list)

	return
}

// 创建图片
func (p *Picture) Create() error {
	return db.Db.Create(&p).Error
}
