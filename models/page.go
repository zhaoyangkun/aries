package models

import (
	"aries/config/db"
	"aries/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

// 页面
type Page struct {
	gorm.Model
	Title  string `gorm:"type:varchar(100);not null;" json:"title"` // 标题
	Url    string `gorm:"type:varchar(100);not null;" json:"url"`   // 访问地址
	Html   string `gorm:"type:Text;not null;" json:"html"`          // 页面内容
	MDHtml string `gorm:"type:MediumText;not null;" json:"md_html"` // markdown渲染后页面内容
}

// 获取所有页面
func (Page) GetAll() (list []Page, err error) {
	err = db.Db.Find(&list).Error

	return
}

//根据 ID获取页面
func (Page) GetById(id uint) (page Page, err error) {
	err = db.Db.Where("`id` = ?", id).First(&page).Error

	return
}

// 根据 URL获取页面
func (Page) GetByUrl(url string) (p Page, err error) {
	err = db.Db.Where("`url` = ?", url).First(&p).Error

	return
}

// 分页获取页面
func (Page) GetByPage(page *utils.Pagination, key string) (list []Page, total uint, err error) {
	query := db.Db.Order("`created_at` desc", true).Find(&list)

	if key != "" {
		query = query.Where("`title` like concat('%',?,'%') or `html` like concat('%',?,'%')", key, key)
	}

	total, err = utils.ToPage(page, query, &list)

	return
}

// 创建页面
func (p *Page) Create() error {
	return db.Db.Create(&p).Error
}

// 更新页面
func (p *Page) Update() error {
	return db.Db.Model(&Page{}).Updates(&p).Error
}

// 批量删除页面
func (Page) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",")

	return db.Db.Where("`id` in (?)", idList).Unscoped().Delete(&Page{}).Error
}
