package forms

import (
	"aries/models"
	"aries/utils"

	"github.com/jinzhu/gorm"
)

// LinkPageForm 友链分页表单
type LinkPageForm struct {
	Key              string `form:"key"`         // 关键词
	CategoryId       uint   `form:"category_id"` // 分类 ID
	utils.Pagination        // 分页结构
}

// LinkAddForm 添加友链表单
type LinkAddForm struct {
	CategoryId uint   `json:"category_id" label:"分类 ID"`
	Name       string `json:"name" binding:"required,max=100" label:"网站名称"`
	Url        string `json:"url" binding:"required,max=255,url" label:"网站地址"`
	Desc       string `json:"desc" binding:"max=255" label:"网站描述"`
	Icon       string `json:"icon" binding:"max=255" label:"图标"`
}

// 修改友链表单
type LinkEditForm struct {
	ID         uint   `json:"ID" binding:"required" label:"ID"`
	CategoryId uint   `json:"category_id" label:"分类 ID"`
	Name       string `json:"name" binding:"required,max=100" label:"网站名称"`
	Url        string `json:"url" binding:"required,max=255,url" label:"网站地址"`
	Desc       string `json:"desc" binding:"max=255" label:"网站描述"`
	Icon       string `json:"icon" binding:"max=255" label:"图标"`
}

// 转换添加友链表单数据到友链实体
func (form LinkAddForm) BindToModel() models.Link {
	return models.Link{
		CategoryId: form.CategoryId,
		Name:       form.Name,
		Url:        form.Url,
		Desc:       form.Desc,
		Icon:       form.Icon,
	}
}

// 转换修改友链表单数据到友链实体
func (form LinkEditForm) BindToModel() models.Link {
	return models.Link{
		Model:      gorm.Model{ID: form.ID},
		CategoryId: form.CategoryId,
		Name:       form.Name,
		Url:        form.Url,
		Desc:       form.Desc,
		Icon:       form.Icon,
	}
}
