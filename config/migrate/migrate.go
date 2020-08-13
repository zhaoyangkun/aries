package migrate

import (
	"aries/config/db"
	"aries/models"
)

func Migrate() {
	// 根据实体结构，反向生成数据表（自动迁移）
	db.Db.AutoMigrate(
		&models.Article{}, &models.Category{}, &models.Comment{},
		&models.Tag{}, &models.Theme{}, &models.User{},
		&models.Link{}, &models.Nav{}, &models.Page{},
		&models.SysSetting{}, &models.SysSettingItem{}, &models.ThemeSetting{},
		&models.Picture{},
	)
}
