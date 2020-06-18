package model

import (
	"github.com/jinzhu/gorm"
)

// 评论结构体
type Comment struct {
	gorm.Model
	User            User     `gorm:"foreignkey:UserId" json:"user"`                    // 发表人
	UserId          uint     `json:"user_id"`                                          // 发表人 ID
	ToUser          User     `gorm:"foreignkey:ToUserId" json:"user"`                  // 回复人
	ToUserId        uint     `json:"to_user_id"`                                       // 回复人 ID
	RootComment     *Comment `gorm:"foreignkey:RootCommentId" json:"root_comment"`     // 根评论
	RootCommentId   uint     `json:"root_comment_id"`                                  // 根评论 ID
	ParentComment   *Comment `gorm:"foreignkey:ParentCommentId" json:"parent_comment"` // 父评论
	ParentCommentId uint     `json:"parent_comment_id"`                                // 父评论 ID
	Content         string   `gorm:"type:Text;not null;" json:"content"`               // 评论内容
	MDContent       string   `gorm:"type:MediumText;not null;" json:"md_content"`      // markdown 渲染后评论内容
	IsRecycled      bool     `gorm:"type:bool;default:false;" json:"is_recycled"`      // 是否加入回收站
	IsChecked       bool     `gorm:"type:bool;default:false" json:"is_checked"`        // 是否通过审核
}
