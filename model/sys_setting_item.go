package model

import (
	"aries/config/db"
	"github.com/jinzhu/gorm"
)

// 系统设置条目结构
type SysSettingItem struct {
	gorm.Model
	SysId uint   `json:"sys_id"`                            // 系统设置 ID
	Key   string `gorm:"varchar(100);not null;" json:"key"` // 键
	Val   string `gorm:"varchar(255);not null;" json:"val"` // 值
}

// 批量创建设置条目
func (SysSettingItem) MultiCreate(itemList []SysSettingItem) error {
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
	// 批量插入
	for _, item := range itemList {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
