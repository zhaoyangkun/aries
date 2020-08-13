package models

import "github.com/jinzhu/gorm"

// 图片
type Picture struct {
	gorm.Model
	StorageType string `gorm:"varchar(20);not null" json:"storage_type"` // 存储类型
	Hash        string `gorm:"varchar(100)" json:"hash"`                 // 标识
	FileName    string `gorm:"varchar(255);not null" json:"file_name"`   // 文件名
	URL         string `gorm:"varchar(255);not null" json:"url"`         // 访问地址
	Size        uint   `gorm:"int(6)" json:"size"`                       // 空间大小（KB）
}
