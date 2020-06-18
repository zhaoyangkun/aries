package migrate

import (
	"aries/config/db"
	"aries/model"
)

func Migrate() {
	// 根据实体结构，反向生成数据表（自动迁移）
	db.Db.AutoMigrate(
		&model.Article{},
		&model.Category{},
		&model.Comment{},
		&model.Tag{},
		&model.Theme{},
		&model.User{},
	)
}
