package model

import (
	"aries/config/db"
	"github.com/jinzhu/gorm"
)

// 评论结构体
type Comment struct {
	gorm.Model
	ParentComment   *Comment `gorm:"foreignkey:ParentCommentId" json:"parent_comment"` // 父评论
	ParentCommentId uint     `json:"parent_comment_id"`                                // 父评论 ID
	RootComment     *Comment `gorm:"foreignkey:RootCommentId" json:"root_comment"`     // 根评论
	RootCommentId   uint     `json:"root_comment_id"`                                  // 根评论 ID
	User            User     `gorm:"foreignkey:UserId" json:"user"`                    // 发表人
	UserId          uint     `json:"user_id"`                                          // 发表人 ID
	Content         string   `gorm:"type:Text;not null;" json:"content"`               // 评论内容
	MDContent       string   `gorm:"type:MediumText;not null;" json:"md_content"`      // markdown 渲染后评论内容
}

// 初始化数据表
func init() {
	db.Db.AutoMigrate(&Comment{})
}
