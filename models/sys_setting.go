package models

import (
	"aries/config/db"

	"github.com/jinzhu/gorm"
)

// SysSetting 系统设置
type SysSetting struct {
	gorm.Model
	Items []SysSettingItem `gorm:"ForeignKey:SysId" json:"items"`          // 设置条目列表
	Name  string           `gorm:"type:varchar(50);not null;" json:"name"` // 名称
}

// GetByName 根据名称获取系统设置
func (SysSetting) GetByName(name string) (s SysSetting, err error) {
	err = db.Db.Where("`name` = ?", name).First(&s).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// Create 创建系统设置
func (s *SysSetting) Create() (err error) {
	err = db.Db.Create(&s).Error
	return
}
