package forms

import (
	"aries/models"
	"aries/utils"

	"github.com/jinzhu/gorm"
)

// CommentPageForm 评论分页表单
type CommentPageForm struct {
	Key              string `form:"key"`        // 关键词
	Type             uint   `form:"type"`       // 类型
	State            uint   `form:"state"`      // 状态
	ArticleId        uint   `form:"article_id"` // 文章 ID
	PageId           uint   `form:"page_id"`    // 页面 ID
	IsParent         uint   `form:"is_parent"`  // 是否为父评论
	utils.Pagination        // 分页结构
}

// CommentAddForm 添加评论表单
type CommentAddForm struct {
	AdminUserId     uint   `json:"admin_user_id" label:"博主 ID"`                            // 博主 ID
	ArticleId       uint   `json:"article_id" label:"文章 ID"`                               // 文章 ID
	PageId          uint   `json:"page_id" label:"页面 ID"`                                  // 页面 ID
	RootCommentId   uint   `json:"root_comment_id" label:"根评论 ID"`                         // 根评论 ID
	ParentCommentId uint   `json:"parent_comment_id" label:"父评论 ID"`                       // 父评论 ID
	Type            uint   `json:"type" binding:"required,min=1,max=4" label:"类型"`         // 类型，1 表文章评论，2 表友链页评论，3 表关于页评论，4 表示自定义页面评论
	Email           string `json:"email" binding:"required,max=50,email" label:"邮箱"`       // 邮箱
	Url             string `json:"url" binding:"required,max=150,url" label:"网站"`          // 访问地址
	NickName        string `json:"nick_name" binding:"required,max=50" label:"呢称"`         // 昵称
	UserImg         string `json:"user_img" binding:"required,max=10000" label:"用户头像"`     // 用户头像
	Device          string `json:"device" binding:"required,max=100" label:"设备"`           // 设备
	Content         string `json:"content" binding:"required,min=6,max=1200" label:"评论内容"` // 评论内容
	MDContent       string `json:"md_content" label:"markdown 渲染后评论内容"`                    // markdown 渲染后评论内容
	IsRecycled      bool   `json:"is_recycled"`                                            // 是否加入回收站
	IsChecked       bool   `json:"is_checked"`                                             // 是否通过审核
}

// CommentEditForm 修改评论表单
type CommentEditForm struct {
	ID              uint   `json:"id" binding:"required" label:"ID"`                       // ID
	AdminUserId     uint   `json:"admin_user_id" label:"博主 ID"`                            // 博主 ID
	ArticleId       uint   `json:"article_id" label:"文章 ID"`                               // 文章 ID
	PageId          uint   `json:"page_id" label:"页面 ID"`                                  // 页面 ID
	RootCommentId   uint   `json:"root_comment_id" label:"根评论 ID"`                         // 根评论 ID
	ParentCommentId uint   `json:"parent_comment_id" label:"父评论 ID"`                       // 父评论 ID
	Type            uint   `json:"type" binding:"required,min=1,max=4" label:"类型"`         // 类型，1 表文章评论，2 表友链页评论，3 表关于页评论，4 表示自定义页面评论
	Email           string `json:"email" binding:"required,max=50,email" label:"邮箱"`       // 邮箱
	Url             string `json:"url" binding:"required,max=150,url" label:"网站"`          // 访问地址
	NickName        string `json:"nick_name" binding:"required,max=50" label:"呢称"`         // 昵称
	Content         string `json:"content" binding:"required,min=6,max=1200" label:"评论内容"` // 评论内容
	MDContent       string `json:"md_content"`                                             // markdown 渲染后评论内容
	Device          string `json:"device" binding:"required,max=100" label:"设备"`           // 设备
	IsRecycled      bool   `json:"is_recycled"`                                            // 是否加入回收站
	IsChecked       bool   `json:"is_checked"`                                             // 是否通过审核
}

// BindToModel 绑定添加评论表单数据到实体结构
func (form CommentAddForm) BindToModel() models.Comment {
	return models.Comment{
		AdminUserId:     form.AdminUserId,
		ArticleId:       form.ArticleId,
		PageId:          form.PageId,
		RootCommentId:   form.RootCommentId,
		ParentCommentId: form.ParentCommentId,
		Type:            form.Type,
		Email:           form.Email,
		Url:             form.Url,
		NickName:        form.NickName,
		UserImg:         form.UserImg,
		Content:         form.Content,
		MDContent:       form.MDContent,
		Device:          form.Device,
		IsRecycled:      form.IsRecycled,
		IsChecked:       form.IsChecked,
	}
}

// BindToModel 绑定修改评论表单数据到实体结构
func (form CommentEditForm) BindToModel() models.Comment {
	return models.Comment{
		Model:           gorm.Model{ID: form.ID},
		AdminUserId:     form.AdminUserId,
		ArticleId:       form.ArticleId,
		PageId:          form.PageId,
		RootCommentId:   form.RootCommentId,
		ParentCommentId: form.ParentCommentId,
		Type:            form.Type,
		Email:           form.Email,
		Url:             form.Url,
		NickName:        form.NickName,
		Content:         form.Content,
		MDContent:       form.MDContent,
		Device:          form.Device,
		IsRecycled:      form.IsRecycled,
		IsChecked:       form.IsChecked,
	}
}
