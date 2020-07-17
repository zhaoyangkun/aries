package model

import (
	"aries/config/db"
	"aries/config/setting"
	"aries/util"
	"github.com/jinzhu/gorm"
	"strings"
)

// 评论结构
type Comment struct {
	gorm.Model
	Article         Article  `gorm:"foreignkey:ArticleId" json:"article"`              // 文章
	ArticleId       uint     `json:"article_id"`                                       // 文章 ID
	AdminUser       User     `gorm:"foreignkey:AdminUserId" json:"admin_user"`         // 博主
	AdminUserId     uint     `json:"admin_user_id"`                                    // 博主 ID
	RootComment     *Comment `gorm:"foreignkey:RootCommentId" json:"root_comment"`     // 根评论
	RootCommentId   uint     `json:"root_comment_id"`                                  // 根评论 ID
	ParentComment   *Comment `gorm:"foreignkey:ParentCommentId" json:"parent_comment"` // 父评论
	ParentCommentId uint     `json:"parent_comment_id"`                                // 父评论 ID
	Email           string   `gorm:"type:varchar(50);not null;" json:"email"`          // 邮箱
	Url             string   `gorm:"varchar(150);not null;" json:"url"`                // 访问地址
	NickName        string   `gorm:"varchar(50);not null;" json:"nick_name"`           // 昵称
	Content         string   `gorm:"type:Text;not null;" json:"content"`               // 评论内容
	MDContent       string   `gorm:"type:MediumText;not null;" json:"md_content"`      // markdown 渲染后评论内容
	IsRecycled      *bool    `gorm:"type:bool;default:false;" json:"is_recycled"`      // 是否加入回收站
	IsChecked       *bool    `gorm:"type:bool;default:false" json:"is_checked"`        // 是否通过审核
}

// 获取所有评论
func (Comment) GetAll() (list []Comment, err error) {
	err = db.Db.Preload("Article").
		Order("created_at desc", true).Find(&list).Error
	return
}

// 分页获取评论
func (Comment) GetByPage(page *util.Pagination, key string, state uint) (list []Comment,
	total uint, err error) {
	query := db.Db.Preload("Article").
		Order("created_at desc", true).Find(&list)
	if key != "" {
		query = query.Where("content like concat('%',?,'%')", key)
	}
	if state > 0 {
		switch state {
		case 1:
			query = query.Where("is_recycled = 1")
		case 2:
			query = query.Where("is_checked = 0 and is_recycled = 0")
		case 3:
			query = query.Where("is_checked = 1 and is_recycled = 0")
		default:
			break
		}
	}
	total, err = util.ToPage(page, query, &list)
	return
}

// 创建评论
func (comment Comment) Create() error {
	comment.MDContent = setting.LuteEngine.MarkdownStr("", comment.Content)
	return db.Db.Create(&comment).Error
}

// 更新评论
func (comment Comment) Update() error {
	comment.MDContent = setting.LuteEngine.MarkdownStr("", comment.Content)
	return db.Db.Model(&Comment{}).Updates(&comment).Error
}

// 删除评论
func (Comment) DeleteById(id string) error {
	return db.Db.Where("id = ?", id).Unscoped().Delete(&Comment{}).Error
}

// 批量删除评论
func (Comment) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",")
	return db.Db.Where("id in (?)", idList).Unscoped().Delete(&Comment{}).Error
}
