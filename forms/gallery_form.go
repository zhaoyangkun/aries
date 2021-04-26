package forms

import (
	"aries/models"
	"aries/utils"

	"github.com/jinzhu/gorm"
)

// GalleryPageForm 图库分页表单
type GalleryPageForm struct {
	CategoryId uint   `form:"category_id"` // 分类 ID
	Key        string `form:"key"`         // 关键词
	utils.Pagination
}

// AddGalleryForm 添加图库表单
type AddGalleryForm struct {
	CategoryId uint   `json:"category_id" label:"分类 ID"`
	URL        string `binding:"required,max=255,url" json:"url" label:"图片地址"`
	Name       string `binding:"required,max=255" json:"name" label:"图片名称"`
	Desc       string `binding:"max=255" json:"desc" label:"图片描述"`
	Location   string `binding:"max=50" json:"location" label:"拍摄地点"`
}

// 修改图库表单
type EditGalleryForm struct {
	ID         uint   `json:"id" binding:"required" label:"ID"`
	CategoryId uint   `json:"category_id" label:"分类 ID"`
	URL        string `binding:"required,max=255,url" json:"url" label:"图片地址"`
	Name       string `binding:"required,max=255" json:"name" label:"图片名称"`
	Desc       string `binding:"max=255" json:"desc" label:"图片描述"`
	Location   string `binding:"max=50" json:"location" label:"拍摄地点"`
}

// 绑定添加图库表单数据到图库实体
func (form AddGalleryForm) BindToModel() models.Gallery {
	return models.Gallery{
		CategoryId: form.CategoryId,
		URL:        form.URL,
		Name:       form.Name,
		Desc:       form.Desc,
		Location:   form.Location,
	}
}

// 绑定修改图库表单数据到图库实体
func (form EditGalleryForm) BindToModel() models.Gallery {
	return models.Gallery{
		Model:      gorm.Model{ID: form.ID},
		CategoryId: form.CategoryId,
		URL:        form.URL,
		Name:       form.Name,
		Desc:       form.Desc,
		Location:   form.Location,
	}
}
