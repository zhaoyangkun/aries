package forms

import (
	"aries/models"
	"aries/utils"

	"github.com/jinzhu/gorm"
)

// TagPageForm 标签分页表单
type TagPageForm struct {
	Key              string `form:"key"` // 关键词
	utils.Pagination        // 分页结构
}

// TagAddForm 添加标签表单
type TagAddForm struct {
	Name string `json:"name" binding:"required,max=100" label:"标签名称"`
}

// TagEditForm 修改标签表单
type TagEditForm struct {
	ID   uint   `json:"ID" binding:"required" label:"ID"`
	Name string `json:"name" binding:"required,max=100" label:"标签名称"`
}

// BindToModel 绑定添加表单数据到实体结构
func (form TagAddForm) BindToModel() models.Tag {
	return models.Tag{
		Name: form.Name,
	}
}

// BindToModel 绑定修改表单数据到实体结构
func (form TagEditForm) BindToModel() models.Tag {
	return models.Tag{
		Model: gorm.Model{ID: form.ID},
		Name:  form.Name,
	}
}
