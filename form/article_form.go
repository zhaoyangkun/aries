package form

import (
	"aries/model"
	"aries/util"
	"github.com/jinzhu/gorm"
)

// 文章分页表单
type ArticlePageForm struct {
	Key             string `form:"key"`         // 关键词
	State           uint   `form:"state"`       // 状态
	CategoryId      uint   `form:"category_id"` // 分类 ID
	util.Pagination        // 分页结构
}

// 添加文章表单
type ArticleAddForm struct {
	UserId           uint   `json:"user_id" binding:"required" label:"用户 ID"`             // 用户 ID
	CategoryId       uint   `json:"category_id" binding:"required" label:"分类 ID"`         // 分类 ID
	OrderId          uint   `json:"order_id" label:"排序 ID"`                               // 排序 ID
	TagIds           string `json:"tag_ids" label:"标签"`                                   // 标签
	IsTop            bool   `json:"is_top" label:"是否置顶"`                                  // 是否置顶
	IsRecycled       bool   `json:"is_recycled" label:"是否回收"`                             // 是否回收
	IsPublished      bool   `json:"is_published" label:"是否发布"`                            // 是否发布
	IsAllowCommented bool   `json:"is_allow_commented" label:"是否允许评论"`                    // 是否允许评论
	Pwd              string `json:"pwd" binding:"max=30" label:"访问密码"`                    // 访问密码
	URL              string `json:"url" binding:"max=255" label:"访问 URL"`                 // 访问 URL
	Title            string `json:"title" binding:"required,max=255" label:"标题"`          // 标题
	Summary          string `json:"summary" binding:"max=255" label:"摘要"`                 // 摘要
	Img              string `json:"img" binding:"max=255" label:"图片"`                     // 图片
	Content          string `json:"content" binding:"required,max=100000" label:"内容"`     // 内容
	MDContent        string `json:"md_content" binding:"required" label:"Markdown 渲染后内容"` // Markdown 渲染后内容
	Keywords         string `json:"keywords" binding:"required" label:"SEO 关键词"`          // SEO 关键词
}

// 修改文章表单
type ArticleEditForm struct {
	ID               uint   `json:"id" binding:"required" label:"ID"`                     // ID
	UserId           uint   `json:"user_id" binding:"required" label:"用户 ID"`             // 用户 ID
	CategoryId       uint   `json:"category_id" binding:"required" label:"分类 ID"`         // 分类 ID
	OrderId          uint   `json:"order_id" label:"排序 ID"`                               // 排序 ID
	TagIds           string `json:"tag_ids" label:"标签"`                                   // 标签
	IsTop            bool   `json:"is_top" label:"是否置顶"`                                  // 是否置顶
	IsRecycled       bool   `json:"is_recycled" label:"是否回收"`                             // 是否回收
	IsPublished      bool   `json:"is_published" label:"是否发布"`                            // 是否发布
	IsAllowCommented bool   `json:"is_allow_commented" label:"是否允许评论"`                    // 是否允许评论
	Pwd              string `json:"pwd" binding:"max=30" label:"访问密码"`                    // 访问密码
	URL              string `json:"url" binding:"max=255" label:"访问 URL"`                 // 访问 URL
	Title            string `json:"title" binding:"required,max=255" label:"标题"`          // 标题
	Summary          string `json:"summary" binding:"max=255" label:"摘要"`                 // 摘要
	Img              string `json:"img" binding:"max=255" label:"图片"`                     // 图片
	Content          string `json:"content" binding:"required,max=100000" label:"内容"`     // 内容
	MDContent        string `json:"md_content" binding:"required" label:"Markdown 渲染后内容"` // Markdown 渲染后内容
	Keywords         string `json:"keywords" binding:"required" label:"SEO 关键词"`          // SEO 关键词
}

// 绑定添加文章表单数据到实体结构
func (form ArticleAddForm) BindToModel() model.Article {
	return model.Article{
		UserId:           form.UserId,
		CategoryId:       form.CategoryId,
		OrderId:          form.OrderId,
		IsTop:            &form.IsTop,
		IsRecycled:       &form.IsRecycled,
		IsPublished:      &form.IsPublished,
		IsAllowCommented: &form.IsAllowCommented,
		Pwd:              form.Pwd,
		URL:              form.URL,
		Title:            form.Title,
		Summary:          form.Summary,
		Img:              form.Img,
		Content:          form.Content,
		MDContent:        form.MDContent,
		Keywords:         form.Keywords,
	}
}

// 绑定修改文章表单数据到实体结构
func (form ArticleEditForm) BindToModel() model.Article {
	return model.Article{
		Model:            gorm.Model{ID: form.ID},
		UserId:           form.UserId,
		CategoryId:       form.CategoryId,
		OrderId:          form.OrderId,
		IsTop:            &form.IsTop,
		IsRecycled:       &form.IsRecycled,
		IsPublished:      &form.IsPublished,
		IsAllowCommented: &form.IsAllowCommented,
		Pwd:              form.Pwd,
		URL:              form.URL,
		Title:            form.Title,
		Summary:          form.Summary,
		Img:              form.Img,
		Content:          form.Content,
		MDContent:        form.MDContent,
		Keywords:         form.Keywords,
	}
}
