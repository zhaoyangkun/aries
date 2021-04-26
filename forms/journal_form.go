package forms

import (
	"aries/models"
	"aries/utils"

	"github.com/jinzhu/gorm"
)

// JournalPageForm 日志分页表单
type JournalPageForm struct {
	Key              string `form:"key"` // 关键词
	utils.Pagination        // 分页结构
}

// JournalAddForm 日志添加表单
type JournalAddForm struct {
	IsSecret bool   `json:"is_secret" label:"是否私密"`
	Content  string `binding:"required,max=255" json:"content" label:"日志内容"`
}

// JournalEditForm 日志修改表单
type JournalEditForm struct {
	ID       uint   `json:"id" binding:"required" label:"ID"`
	IsSecret bool   `json:"is_secret" label:"是否私密"`
	Content  string `binding:"required,max=255" json:"content" label:"日志内容"`
}

// BindToModel 转换日志添加表单数据 -> 日志实体
func (form JournalAddForm) BindToModel() models.Journal {
	return models.Journal{
		IsSecret: form.IsSecret,
		Content:  form.Content,
	}
}

// BindToModel 转换日志修改表单数据 -> 日志实体
func (form JournalEditForm) BindToModel() models.Journal {
	return models.Journal{
		Model:    gorm.Model{ID: form.ID},
		IsSecret: form.IsSecret,
		Content:  form.Content,
	}
}
