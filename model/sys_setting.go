package model

import "github.com/jinzhu/gorm"

// 系统设置结构
type SysSetting struct {
	gorm.Model
	Items []SysSettingItem `gorm:"foreignkey:SysId" json:"items"`     // 设置条目列表
	Name  string           `gorm:"varchar(50);not null;" json:"name"` // 名称
}
