package forms

import (
	"aries/models"
	"aries/utils"

	"github.com/jinzhu/gorm"
)

// CategoryPageForm 分类分页表单
type CategoryPageForm struct {
	Key              string `form:"key"`                                           // 关键词
	CategoryType     *uint  `form:"category_type" binding:"required" label:"分类类型"` // 分类类型
	utils.Pagination        // 分页结构
}

// ArticleCategoryAddForm 添加文章分类表单
type ArticleCategoryAddForm struct {
	Type     uint   `json:"type"`                                  // 分类类型，默认值为 0 表文章；1 表友链; 2 表示图库
	Name     string `json:"name" binding:"required" label:"分类名称"`  // 分类名称
	Url      string `json:"url" binding:"required" label:"访问 URL"` // 访问 URL
	ParentId uint   `json:"parent_id" label:"父级分类"`                // 父级分类 ID
}

// ArticleCategoryEditForm 修改文章分类表单
type ArticleCategoryEditForm struct {
	ID       uint   `json:"ID" binding:"required" label:"ID"`
	Type     uint   `json:"type"`                                  // 分类类型，默认值为 0 表文章；1 表友链; 2 表示图库
	Name     string `json:"name" binding:"required" label:"分类名称"`  // 分类名称
	Url      string `json:"url" binding:"required" label:"访问 URL"` // 访问 URL
	ParentId uint   `json:"parent_id" label:"父级分类"`                // 父级分类 ID
}

// LinkCategoryAddForm 添加友链分类表单
type LinkCategoryAddForm struct {
	Type uint   `json:"type"`                                 // 分类类型，默认值为 0 表文章；1 表友链; 2 表示图库
	Name string `json:"name" binding:"required" label:"分类名称"` // 分类名称
}

// LinkCategoryEditForm 修改友链分类表单
type LinkCategoryEditForm struct {
	ID   uint   `json:"ID" binding:"required" label:"ID"`
	Type uint   `json:"type"`                                 // 分类类型，默认值为 0 表文章；1 表友链; 2 表示图库
	Name string `json:"name" binding:"required" label:"分类名称"` // 分类名称
}

// GalleryCategoryAddForm 添加图库分类表单
type GalleryCategoryAddForm struct {
	Type uint   `json:"type"`                                 // 分类类型，默认值为 0 表文章；1 表友链; 2 表示图库
	Name string `json:"name" binding:"required" label:"分类名称"` // 分类名称
}

// GalleryCategoryEditForm 修改图库分类表单
type GalleryCategoryEditForm struct {
	ID   uint   `json:"ID" binding:"required" label:"ID"`
	Type uint   `json:"type"`                                 // 分类类型，默认值为 0 表文章；1 表友链; 2 表示图库
	Name string `json:"name" binding:"required" label:"分类名称"` // 分类名称
}

// BindToModel 绑定添加文章分类表单到分类结构
func (form ArticleCategoryAddForm) BindToModel() models.Category {
	return models.Category{
		Type:     form.Type,
		Name:     form.Name,
		Url:      form.Url,
		ParentId: form.ParentId,
	}
}

// BindToModel 绑定修改文章分类表单到分类结构
func (form ArticleCategoryEditForm) BindToModel() models.Category {
	return models.Category{
		Model:    gorm.Model{ID: form.ID},
		Type:     form.Type,
		Name:     form.Name,
		Url:      form.Url,
		ParentId: form.ParentId,
	}
}

// BindToModel 绑定添加友链分类表单数据到实体
func (form LinkCategoryAddForm) BindToModel() models.Category {
	return models.Category{
		Type: form.Type,
		Name: form.Name,
	}
}

// 绑定修改友链分类表单数据到实体
func (form LinkCategoryEditForm) BindToModel() models.Category {
	return models.Category{
		Model: gorm.Model{ID: form.ID},
		Type:  form.Type,
		Name:  form.Name,
	}
}

// 绑定添加图库分类表单数据到实体
func (form GalleryCategoryAddForm) BindToModel() models.Category {
	return models.Category{
		Type: form.Type,
		Name: form.Name,
	}
}

// 绑定修改图库分类表单数据到实体
func (form GalleryCategoryEditForm) BindToModel() models.Category {
	return models.Category{
		Model: gorm.Model{ID: form.ID},
		Type:  form.Type,
		Name:  form.Name,
	}
}
