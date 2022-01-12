package models

import (
	"aries/config/db"
	"aries/config/setting"
	"aries/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

// Page 页面
type Page struct {
	gorm.Model
	Title  string `gorm:"type:varchar(100);not null;" json:"title"` // 标题
	Url    string `gorm:"type:varchar(100);not null;" json:"url"`   // 访问地址
	Html   string `gorm:"type:Text;not null;" json:"html"`          // 页面内容
	MDHtml string `gorm:"type:MediumText;not null;" json:"md_html"` // markdown 渲染后页面内容
}

// GetAll 获取所有页面
func (Page) GetAll() (list []Page, err error) {
	err = db.Db.Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetById 根据 ID 获取页面
func (Page) GetById(id uint) (page Page, err error) {
	err = db.Db.Where("`id` = ?", id).First(&page).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByUrl 根据 URL 获取页面
func (Page) GetByUrl(url string) (p Page, err error) {
	err = db.Db.Where("`url` = ?", url).First(&p).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByPage 分页获取页面
func (Page) GetByPage(page *utils.Pagination, key string) (list []Page, total uint, err error) {
	query := db.Db.Order("`created_at` desc", true).Find(&list)

	if key != "" {
		query = query.Where("`title` like concat('%',?,'%') or `html` like concat('%',?,'%')", key, key)
	}

	total, err = utils.ToPage(page, query, &list)
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// Create 创建页面
func (p *Page) Create() error {
	p.MDHtml = setting.LuteEngine.MarkdownStr("", p.Html)
	return db.Db.Create(&p).Error
}

// Update 更新页面
func (p *Page) Update() error {
	p.MDHtml = setting.LuteEngine.MarkdownStr("", p.Html)
	return db.Db.Model(&Page{}).Updates(&p).Error
}

// MultiDelByIds 批量删除页面
func (Page) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",")

	return db.Db.Where("`id` in (?)", idList).Unscoped().Delete(&Page{}).Error
}
