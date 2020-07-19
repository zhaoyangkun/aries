package form

import (
	"aries/model"
	"aries/util"
	"github.com/jinzhu/gorm"
)

// 评论分页表单
type CommentPageForm struct {
	Key             string `form:"key"`   // 关键词
	State           uint   `form:"state"` // 状态
	util.Pagination        // 分页结构
}

// 添加评论表单
type CommentAddForm struct {
	ArticleId       uint   `json:"article_id" binding:"required" label:"文章 ID"`            // 文章 ID
	AdminUserId     uint   `json:"admin_user_id"`                                          // 博主 ID
	RootCommentId   uint   `json:"root_comment_id"`                                        // 根评论 ID
	ParentCommentId uint   `json:"parent_comment_id"`                                      // 父评论 ID
	Email           string `json:"email" binding:"required,max=50,email" label:"邮箱"`       // 邮箱
	Url             string `json:"url" binding:"required,max=150,url" label:"网站"`          // 访问地址
	NickName        string `json:"nick_name" binding:"required,max=50" label:"呢称"`         // 昵称
	Content         string `json:"content" binding:"required,min=6,max=1200" label:"评论内容"` // 评论内容
	MDContent       string `json:"md_content"`                                             // markdown 渲染后评论内容
	IsRecycled      bool   `json:"is_recycled"`                                            // 是否加入回收站
	IsChecked       bool   `json:"is_checked"`                                             // 是否通过审核
}

// 修改评论表单
type CommentEditForm struct {
	ID              uint   `json:"id" binding:"required" label:"ID"`                       // ID
	ArticleId       uint   `json:"article_id" binding:"required" label:"文章 ID"`            // 文章 ID
	AdminUserId     uint   `json:"admin_user_id"`                                          // 博主 ID
	RootCommentId   uint   `json:"root_comment_id"`                                        // 根评论 ID
	ParentCommentId uint   `json:"parent_comment_id"`                                      // 父评论 ID
	Email           string `json:"email" binding:"required,max=50,email" label:"邮箱"`       // 邮箱
	Url             string `json:"url" binding:"required,max=150,url" label:"网站"`          // 访问地址
	NickName        string `json:"nick_name" binding:"required,max=50" label:"呢称"`         // 昵称
	Content         string `json:"content" binding:"required,min=6,max=1200" label:"评论内容"` // 评论内容
	MDContent       string `json:"md_content"`                                             // markdown 渲染后评论内容
	IsRecycled      bool   `json:"is_recycled"`                                            // 是否加入回收站
	IsChecked       bool   `json:"is_checked"`                                             // 是否通过审核
}

// 绑定添加评论表单数据到实体结构
func (form CommentAddForm) BindToModel() model.Comment {
	return model.Comment{
		ArticleId:       form.ArticleId,
		AdminUserId:     form.AdminUserId,
		RootCommentId:   form.RootCommentId,
		ParentCommentId: form.ParentCommentId,
		Email:           form.Email,
		Url:             form.Url,
		NickName:        form.NickName,
		Content:         form.Content,
		MDContent:       form.MDContent,
		IsRecycled:      &form.IsRecycled,
		IsChecked:       &form.IsChecked,
	}
}

// 绑定修改评论表单数据到实体结构
func (form CommentEditForm) BindToModel() model.Comment {
	return model.Comment{
		Model:           gorm.Model{ID: form.ID},
		ArticleId:       form.ArticleId,
		AdminUserId:     form.AdminUserId,
		RootCommentId:   form.RootCommentId,
		ParentCommentId: form.ParentCommentId,
		Email:           form.Email,
		Url:             form.Url,
		NickName:        form.NickName,
		Content:         form.Content,
		MDContent:       form.MDContent,
		IsRecycled:      &form.IsRecycled,
		IsChecked:       &form.IsChecked,
	}
}
