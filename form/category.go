package form

import (
	"aries/model"
	"aries/util"
	"github.com/jinzhu/gorm"
)

// 分类分页表单
type CategoryPageForm struct {
	Key             string `form:"key"` // 关键词
	util.Pagination        // 分页结构
}

// 添加分类表单
type CategoryAddForm struct {
	Type             uint   `json:"type"`                    // 分类类型，默认值为 0 表文章；1 表友链
	Name             string `json:"name" binding:"required"` // 分类名称
	Url              string `json:"url" binding:"required"`  // 访问 URL
	ParentCategoryId uint   `json:"parent_category_id"`      // 父级分类 ID
}

// 修改分类表单
type CategoryEditForm struct {
	Id               uint
	Type             uint   `json:"type"`                    // 分类类型，默认值为 0 表文章；1 表友链
	Name             string `json:"name" binding:"required"` // 分类名称
	Url              string `json:"url" binding:"required"`  // 访问 URL
	ParentCategoryId uint   `json:"parent_category_id"`      // 父级分类 ID
}

// 绑定添加分类表单到分类结构
func (form CategoryAddForm) BindToModel() model.Category {
	return model.Category{
		Type:             form.Type,
		Name:             form.Name,
		Url:              form.Url,
		ParentCategoryId: form.ParentCategoryId,
	}
}

// 绑定添加分类表单到分类结构
func (form CategoryEditForm) BindToModel() model.Category {
	return model.Category{
		Model:            gorm.Model{ID: form.Id},
		Type:             form.Type,
		Name:             form.Name,
		Url:              form.Url,
		ParentCategoryId: form.ParentCategoryId,
	}
}
