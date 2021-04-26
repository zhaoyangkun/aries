package forms

import (
	"aries/models"

	"github.com/jinzhu/gorm"
)

// NavAddForm 添加菜单表单
type NavAddForm struct {
	ParentNavId uint   `json:"parent_nav_id" label:"父级菜单 ID"`
	OpenType    uint   `json:"open_type" binding:"max=1" label:"打开方式"`
	Name        string `json:"name" binding:"required,max=100" label:"名称"`
	Url         string `json:"url" binding:"required,max=255" label:"访问地址"`
	Icon        string `json:"icon" binding:"max=255" label:"图标"`
}

// NavEditForm 修改菜单表单
type NavEditForm struct {
	ID          uint   `json:"id" binding:"required" label:"ID"`
	ParentNavId uint   `json:"parent_nav_id" label:"父级菜单 ID"`
	OpenType    uint   `json:"open_type" binding:"max=1" label:"打开方式"`
	Name        string `json:"name" binding:"required,max=100" label:"名称"`
	Url         string `json:"url" binding:"required,max=255" label:"访问地址"`
	Icon        string `json:"icon" binding:"max=255" label:"图标"`
}

// BindToModel 绑定添加菜单表单到菜单实体
func (form NavAddForm) BindToModel() models.Nav {
	return models.Nav{
		ParentNavId: form.ParentNavId,
		OpenType:    form.OpenType,
		Name:        form.Name,
		Url:         form.Url,
		Icon:        form.Icon,
	}
}

// BindToModel 绑定修改菜单表单到菜单实体
func (form NavEditForm) BindToModel() models.Nav {
	return models.Nav{
		Model:       gorm.Model{ID: form.ID},
		ParentNavId: form.ParentNavId,
		OpenType:    form.OpenType,
		Name:        form.Name,
		Url:         form.Url,
		Icon:        form.Icon,
	}
}
