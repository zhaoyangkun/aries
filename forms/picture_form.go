package forms

import (
	"aries/models"
	"aries/utils"
)

// PicturePageForm 图片分页表单
type PicturePageForm struct {
	Key              string `form:"key"`          // 关键词
	StorageName      string `form:"storage_name"` //存储类型
	utils.Pagination        // 分页结构
}

// PictureAddForm 添加图片表单
type PictureAddForm struct {
	StorageType string `json:"storage_type" binding:"required,max=20" label:"存储类型"`
	Hash        string `json:"hash" binding:"max=100" label:"Hash"`
	FileName    string `json:"file_name" binding:"required,max=255" label:"图片名称"`
	URL         string `json:"url" binding:"required,max=255" label:"图片访问地址"`
	Size        uint   `json:"size" binding:"required" label:"图片大小"`
}

// BindToModel 绑定添加图片表单数据到图片实体类
func (form PictureAddForm) BindToModel() models.Picture {
	return models.Picture{
		StorageType: form.StorageType,
		Hash:        form.Hash,
		FileName:    form.FileName,
		URL:         form.URL,
		Size:        form.Size,
	}
}
