package model

import (
	"github.com/jinzhu/gorm"
)

// 文章结构体
type Article struct {
	gorm.Model
	User             User     `gorm:"foreignkey:UserId;not null;" json:"user"`      // 用户
	UserId           uint     `json:"user_id"`                                      // 用户 ID
	Category         Category `gorm:"foreignKey:CategoryId;not null;"`              // 分类
	CategoryId       uint     `json:"category_id"`                                  // 分类 ID
	TagList          []Tag    `gorm:"many2many:tag_article" json:"tag_list"`        // 标签列表
	IsTop            bool     `gorm:"type:bool;default:false;" json:"is_top"`       // 是否置顶
	IsRecycled       bool     `gorm:"type:bool;default:false;" json:"is_recycled"`  // 是否回收
	IsPublished      bool     `gorm:"type:bool;default:true;" json:"is_published"`  // 是否发布
	IsAllowCommented bool     `gorm:"type:bool;default:true;" json:"allow_comment"` // 是否允许评论
	URL              string   `gorm:"type:varchar(255);not null;" json:"title"`     // 访问 URL
	Title            string   `gorm:"type:varchar(255);not null;" json:"title"`     // 标题
	Summary          string   `gorm:"type:varchar(255);not null;" json:"summary"`   // 摘要
	Img              string   `gorm:"type:varchar(255);not null;" json:"img"`       // 图片
	Content          string   `gorm:"type:Text;not null;" json:"content"`           // 内容
	MDContent        string   `gorm:"type:MediumText;not null;" json:"md_content"`  // Markdown 渲染后内容
	Keywords         string   `gorm:"type:varchar(255);not null;" json:"keywords"`  // SEO 关键词
}
