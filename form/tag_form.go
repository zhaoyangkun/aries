package form

import (
	"aries/model"
	"aries/util"
	"github.com/jinzhu/gorm"
)

// 标签分页表单
type TagPageForm struct {
	Key             string `form:"key"` // 关键词
	util.Pagination        // 分页结构
}

// 添加标签表单
type TagAddForm struct {
	Name string `json:"name" binding:"required,max=100" label:"标签名称"`
}

// 修改标签表单
type TagEditForm struct {
	ID   uint   `json:"id" binding:"required" label:"ID"`
	Name string `json:"name" binding:"required,max=100" label:"标签名称"`
}

// 绑定添加表单数据到实体结构
func (form TagAddForm) BindToModel() model.Tag {
	return model.Tag{
		Name: form.Name,
	}
}

// 绑定修改表单数据到实体结构
func (form TagEditForm) BindToModel() model.Tag {
	return model.Tag{
		Model: gorm.Model{ID: form.ID},
		Name:  form.Name,
	}
}
