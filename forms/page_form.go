package forms

import (
	"aries/models"
	"aries/utils"

	"github.com/jinzhu/gorm"
)

type PageForm struct {
	Key string `form:"key"`
	utils.Pagination
}

type AddPageForm struct {
	Title  string `binding:"required,max=100" json:"title" label:"页面标题"`
	Url    string `binding:"required,max=100" json:"url" label:"访问 URL"`
	Html   string `binding:"required,max=100000" json:"html" label:"页面内容"`
	MDHtml string `binding:"required,max=1000000" json:"md_html" label:"markdown 渲染页面内容"`
}

type EditPageForm struct {
	ID     uint   `binding:"required" json:"id" label:"ID"`
	Title  string `binding:"required,max=100" json:"title" label:"页面标题"`
	Url    string `binding:"required,max=100" json:"url" label:"访问 URL"`
	Html   string `binding:"required,max=100000" json:"html" label:"页面内容"`
	MDHtml string `binding:"required,max=1000000" json:"md_html" label:"markdown 渲染页面内容"`
}

func (form AddPageForm) BindToModel() models.Page {
	return models.Page{
		Title:  form.Title,
		Url:    form.Url,
		Html:   form.Html,
		MDHtml: form.MDHtml,
	}
}

func (form EditPageForm) BindToModel() models.Page {
	return models.Page{
		Model:  gorm.Model{ID: form.ID},
		Title:  form.Title,
		Url:    form.Url,
		Html:   form.Html,
		MDHtml: form.MDHtml,
	}
}
