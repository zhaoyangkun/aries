package models

import (
	"aries/config/db"
	"strconv"

	"github.com/jinzhu/gorm"
)

// SysSettingItem 系统设置条目
type SysSettingItem struct {
	gorm.Model
	SysId uint   `json:"sys_id"`                                 // 系统设置 ID
	Key   string `gorm:"type:varchar(100);not null;" json:"key"` // 键
	Val   string `gorm:"type:Text;not null;" json:"val"`         // 值
}

// GetBySysSettingName 根据设置名称获取系统设置条目
func (SysSettingItem) GetBySysSettingName(settingName string) (map[string]string, error) {
	var sysSetting SysSetting
	var itemList []SysSettingItem
	var err error
	result := map[string]string{}

	if settingName != "" {
		err = db.Db.Where("`name` = ?", settingName).First(&sysSetting).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return result, err
		}
		err = db.Db.Where("`sys_id` = ?", sysSetting.ID).Find(&itemList).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return result, err
		}
	} else {
		err = db.Db.Find(&itemList).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return result, err
		}
	}

	for _, item := range itemList {
		result[item.Key] = item.Val
	}

	if len(itemList) > 0 {
		result["sys_id"] = strconv.Itoa(int(sysSetting.ID))
	}

	return result, err
}

// MultiCreateOrUpdate 批量创建设置条目
func (SysSettingItem) MultiCreateOrUpdate(itemList []SysSettingItem) error {
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

	count := 0
	for i := range itemList {
		count = 0
		err := tx.Model(&SysSettingItem{}).Where("`sys_id` = ? and `key` = ?", itemList[i].SysId, itemList[i].Key).
			Count(&count).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		if count == 0 {
			err = tx.Create(&itemList[i]).Error
		} else {
			err = tx.Model(&SysSettingItem{}).Where("`sys_id` = ? and `key` = ?", itemList[i].SysId, itemList[i].Key).
				Update("val", itemList[i].Val).Error
		}

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
