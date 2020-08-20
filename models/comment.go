package models

import (
	"aries/config/db"
	"aries/config/setting"
	"aries/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

// 评论
type Comment struct {
	gorm.Model
	Page            Page     `gorm:"ForeignKey:PageId" json:"page"`                    // 页面
	PageId          uint     `json:"page_id"`                                          // 页面 ID
	Article         Article  `gorm:"ForeignKey:ArticleId" json:"article"`              // 文章
	ArticleId       uint     `json:"article_id"`                                       // 文章 ID
	AdminUser       User     `gorm:"ForeignKey:AdminUserId" json:"admin_user"`         // 博主
	AdminUserId     uint     `json:"admin_user_id"`                                    // 博主 ID
	RootComment     *Comment `gorm:"ForeignKey:RootCommentId" json:"root_comment"`     // 根评论
	RootCommentId   uint     `json:"root_comment_id"`                                  // 根评论 ID
	ParentComment   *Comment `gorm:"ForeignKey:ParentCommentId" json:"parent_comment"` // 父评论
	ParentCommentId uint     `json:"parent_comment_id"`                                // 父评论 ID
	Type            uint     `gorm:"type:tinyint(1);unsigned;default:1" json:"type"`   // 类型，1 表文章评论，2 表友链页评论，3 表关于页评论，4 表示自定义页面评论
	Email           string   `gorm:"type:varchar(50);not null;" json:"email"`          // 邮箱
	Url             string   `gorm:"varchar(150);not null;" json:"url"`                // 访问地址
	NickName        string   `gorm:"varchar(50);not null;" json:"nick_name"`           // 昵称
	Content         string   `gorm:"type:Text;not null;" json:"content"`               // 评论内容
	MDContent       string   `gorm:"type:MediumText;not null;" json:"md_content"`      // markdown 渲染后评论内容
	IsRecycled      *bool    `gorm:"type:bool;default:false;" json:"is_recycled"`      // 是否加入回收站
	IsChecked       *bool    `gorm:"type:bool;default:false" json:"is_checked"`        // 是否通过审核
}

// 获取评论数量
func (Comment) GetCount() (int, error) {
	count := 0
	err := db.Db.Model(&Comment{}).Count(&count).Error
	return count, err
}

// 获取最近发表的评论
func (Comment) GetLatest(limit uint) (list []Comment, err error) {
	err = db.Db.Preload("Article").Order("created_at desc", true).
		Limit(limit).Find(&list).Error
	return
}

// 获取所有评论
func (Comment) GetAll() (list []Comment, err error) {
	err = db.Db.Preload("Article").
		Order("created_at desc", true).Find(&list).Error
	return
}

// 分页获取评论
func (Comment) GetByPage(page *utils.Pagination, key string, commentType uint, state uint) (list []Comment,
	total uint, err error) {
	query := db.Db.Preload("Article").
		Order("created_at desc", true).Find(&list)

	if key != "" {
		query = query.Where("content like concat('%',?,'%')", key)
	}

	if commentType > 0 {
		switch commentType {
		case 1:
			query = query.Where("type = 1")
		case 2:
			query = query.Where("type > 1")
		default:
			break
		}
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

	total, err = utils.ToPage(page, query, &list)

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
	return db.Db.Where("`id` = ?", id).Unscoped().Delete(&Comment{}).Error
}

// 批量删除评论
func (Comment) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",")
	return db.Db.Where("`id` in (?)", idList).Unscoped().Delete(&Comment{}).Error
}
