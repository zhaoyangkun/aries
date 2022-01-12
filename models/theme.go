package models

import (
	"aries/config/db"
	logger "aries/log"
	"aries/utils"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Theme 主题
type Theme struct {
	gorm.Model
	ThemeInfo
	IsUsed bool `gorm:"type:bool;default:false;" json:"is_used"`
}

// ThemeInfo 主题信息
type ThemeInfo struct {
	AuthorInfo `yaml:"author"`
	ThemeID    string `gorm:"type:varchar(30);not null;" yaml:"id" json:"theme_id"`
	ThemeName  string `gorm:"type:varchar(100);not null;" yaml:"name" json:"theme_name"`
	Desc       string `gorm:"type:varchar(255);" yaml:"desc" json:"desc"`
	Image      string `gorm:"type:varchar(255);not null;" yaml:"image" json:"image"`
	Repo       string `gorm:"type:varchar(255);" yaml:"repo" json:"repo"`
	Version    string `gorm:"type:varchar(30);not null;" yaml:"version" json:"version"`
}

// AuthorInfo 作者信息
type AuthorInfo struct {
	AuthorName string `gorm:"type:varchar(30);not null;" yaml:"name" json:"author_name"`
	Website    string `gorm:"type:varchar(255);" yaml:"website" json:"website"`
	HeadImg    string `gorm:"type:varchar(255);" yaml:"head_img" json:"head_img"`
}

// GetAll 获取所有主题
func (Theme) GetAll() (list []Theme, err error) {
	err = db.Db.Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByThemeName 根据主题名称获取主题
func (Theme) GetByThemeName(themeName string) (obj Theme, err error) {
	err = db.Db.Where("`theme_name` = ?", themeName).Find(&obj).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// ReadThemeInfo 读取主题配置文件
func (t *ThemeInfo) ReadThemeInfo() {
	rootPath, _ := os.Getwd()
	dirName := filepath.Join(rootPath, "resources", "themes")
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		logger.Logger.Sugar().Error("err: ", err.Error())
	}

	yamlPath := ""
	themeInfo := &ThemeInfo{}
	for _, file := range fileInfos {
		yamlPath = filepath.Join(dirName, file.Name(), "theme.yaml")
		if utils.FileIsExists(yamlPath) {
			yamlFile, err := ioutil.ReadFile(yamlPath)
			if err != nil {
				logger.Logger.Sugar().Panic("读取主题配置文件失败：", err.Error())
			}

			err = yaml.Unmarshal(yamlFile, themeInfo)
			if err != nil {
				logger.Logger.Sugar().Panic("主题配置参数转换失败：", err.Error())
			}

			theme := Theme{ThemeInfo: *themeInfo, IsUsed: false}
			err = theme.CreateOrUpdate()
			if err != nil {
				logger.Logger.Sugar().Error("error：", err.Error())
			}
		}
	}
}

// CreateOrUpdate 创建或更新主题
func (t *Theme) CreateOrUpdate() (err error) {
	existTheme := Theme{}
	err = db.Db.Where("`theme_id` = ?", t.ThemeID).First(&existTheme).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}

	if existTheme.ThemeID != "" { // 已存在该主题，仅更新
		return db.Db.Model(&Theme{}).Where("`theme_id` = ?", t.ThemeID).Updates(&t).Error
	}
	return db.Db.Create(&t).Error // 不存在该主题，则创建
}

// EnableTheme 启用主题
func (t Theme) EnableTheme() error {
	// 开始事务
	tx := db.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	// 启用当前主题
	err := tx.Model(&Theme{}).Where("`theme_name` = ?", t.ThemeName).Updates(
		map[string]interface{}{
			"is_used": true,
		},
	).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 将其他主题切换为非启用状态
	err = tx.Model(&Theme{}).Where("`theme_name` != ?", t.ThemeName).Updates(
		map[string]interface{}{
			"is_used": false,
		},
	).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
