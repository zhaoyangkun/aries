package model

import (
	"aries/config/db"
	"github.com/jinzhu/gorm"
)

// 系统设置结构
type SysSetting struct {
	gorm.Model
	Items []SysSettingItem `gorm:"ForeignKey:SysId" json:"items"`     // 设置条目列表
	Name  string           `gorm:"varchar(50);not null;" json:"name"` // 名称
}

// 创建系统设置
func (s *SysSetting) Create() (err error) {
	err = db.Db.Create(s).Error
	return
}
