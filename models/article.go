package models

import (
	"aries/config/db"
	"aries/config/setting"
	"aries/utils"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Article 文章
type Article struct {
	gorm.Model
	User             User     `gorm:"ForeignKey:UserId;not null;" json:"user"`             // 用户
	UserId           uint     `json:"user_id"`                                             // 用户 ID
	Category         Category `gorm:"ForeignKey:CategoryId" json:"category"`               // 分类
	CategoryId       uint     `json:"category_id"`                                         // 分类 ID
	OrderId          uint     `gorm:"type:int;default:0;" json:"order_id"`                 // 排序 ID
	TagList          []Tag    `gorm:"many2many:tag_article" json:"tag_list"`               // 标签列表
	IsTop            bool     `gorm:"type:bool;default:false;" json:"is_top"`              // 是否置顶
	IsRecycled       bool     `gorm:"type:bool;default:false;" json:"is_recycled"`         // 是否回收
	IsPublished      bool     `gorm:"type:bool;default:true;" json:"is_published"`         // 是否发布
	IsAllowCommented bool     `gorm:"type:bool;default:true;" json:"is_allow_commented"`   // 是否允许评论
	Pwd              string   `gorm:"type:varchar(100);" json:"pwd"`                       // 访问密码
	URL              string   `gorm:"type:varchar(255);not null;unique_index;" json:"url"` // 访问 URL
	Title            string   `gorm:"type:varchar(255);not null;" json:"title"`            // 标题
	Summary          string   `gorm:"type:varchar(255);not null;" json:"summary"`          // 摘要
	Img              string   `gorm:"type:varchar(255);not null;" json:"img"`              // 图片
	Content          string   `gorm:"type:MediumText;not null;" json:"content"`            // 内容
	MDContent        string   `gorm:"type:MediumText;not null;" json:"md_content"`         // Markdown 渲染后内容
	Keywords         string   `gorm:"type:varchar(255)" json:"keywords"`                   // SEO 关键词
	CommentCount     uint     `gorm:"type:int;default:0;" json:"comment_count"`            // 评论数
	VisitCount       uint     `gorm:"type:int;default:0;" json:"visit_count"`              // 浏览数
}

// Archive 归档
type Archive struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Count int `json:"count"`
}

// GetCount 获取文章总数
func (Article) GetCount() (int, error) {
	count := 0
	err := db.Db.Model(&Article{}).Count(&count).Error

	return count, err
}

// GetLatest 获取最近发布的文章
func (Article) GetLatest(limit uint) (list []Article, err error) {
	err = db.Db.Order("created_at desc", true).Limit(limit).
		Where("is_published = 1 and is_recycled = 0").Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetAll 获取所有文章
func (Article) GetAll() (list []Article, err error) {
	err = db.Db.Preload("Category").Preload("TagList").
		Order("created_at desc", true).
		Where("is_published = 1 and is_recycled = 0").Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetById 根据 ID 获取文章
func (Article) GetById(id string) (article Article, err error) {
	err = db.Db.Preload("Category").Preload("TagList").
		Where("`id` = ?", id).First(&article).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByCategoryUrl 根据分类名获取文章
func (Article) GetByCategoryUrl(page *utils.Pagination, url string) (list []Article, name string, total uint, err error) {
	category, err := Category{}.GetByUrl(url)
	if err != nil {
		return
	}

	name = category.Name

	query := db.Db.Model(&Article{}).
		Where("is_published = 1 and is_recycled = 0 and category_id = ?", category.ID)
	total, err = utils.ToPage(page, query, &list)
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByTagName 根据标签名获取文章
func (Article) GetByTagName(page *utils.Pagination, tagName string) (list []Article, total uint, err error) {
	limit := page.Size
	offset := (page.Page - 1) * page.Size

	tag, err := Tag{}.GetByName(tagName)
	if err != nil {
		return
	}

	err = db.Db.Raw("select count(`id`) from `articles` where `id` in "+
		"(select `article_id` from `tag_article` "+
		"where is_published = 1 and is_recycled = 0 and `tag_id` = ?)", tag.ID).Row().Scan(&total)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}

	err = db.Db.Raw("select * from `articles` where `id` in "+
		"(select `article_id` from `tag_article` "+
		"where is_published = 1 and is_recycled = 0 and `tag_id` = ?) "+
		"limit ? offset ?", tag.ID, limit, offset).Scan(&list).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}

	return
}

// GetPrevious 获取上一篇文章
func (Article) GetPrevious(orderId uint, isTop bool, ignore bool) (article Article, err error) {
	if ignore {
		if isTop {
			err = db.Db.Raw("select * from `articles`"+
				" where is_published = 1 and is_recycled = 0 and `order_id` < ? and is_top = 1"+
				" order by `order_id` desc limit 1", orderId).Scan(&article).Error
		} else {
			err = db.Db.Raw("select * from `articles`"+
				" where is_published = 1 and is_recycled = 0 and `order_id` < ? and is_top = 0"+
				" order by `order_id` desc limit 1", orderId).Scan(&article).Error
			if article.Title == "" {
				err = db.Db.Raw("select * from `articles`" +
					" where is_published = 1 and is_recycled = 0 and is_top = 1" +
					" order by `order_id` desc limit 1").Scan(&article).Error
			}
		}
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return
		}

		return
	}

	if isTop {
		err = db.Db.Raw("select * from `articles` "+
			" where is_published = 1 and is_recycled = 0 and `order_id` < ? and is_top = 1"+
			" order by `order_id` desc limit 1", orderId).Scan(&article).Error
	} else {
		err = db.Db.Raw("select * from `articles`"+
			" where is_published = 1 and is_recycled = 0 and `order_id` < ? and is_top = 0"+
			" order by `order_id` desc limit 1", orderId).Scan(&article).Error
	}
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}

	return
}

// GetNext 获取下一篇文章
func (Article) GetNext(orderId uint, isTop bool, ignore bool) (article Article, err error) {
	if ignore {
		if isTop {
			err = db.Db.Raw("select * from `articles`"+
				" where is_published = 1 and is_recycled = 0 and `order_id` > ? and is_top = 1"+
				" order by `order_id` asc limit 1", orderId).Scan(&article).Error
			if article.Title == "" {
				err = db.Db.Raw("select * from `articles`" +
					" where is_published = 1 and is_recycled = 0 and is_top = 0" +
					" order by `order_id` asc limit 1").Scan(&article).Error
			}
		} else {
			err = db.Db.Raw("select * from `articles`"+
				" where is_published = 1 and is_recycled = 0 and `order_id` > ? and is_top = 0"+
				" order by `order_id` asc limit 1", orderId).Scan(&article).Error
		}
		return
	}

	if isTop {
		err = db.Db.Raw("select * from `articles`"+
			" where is_published = 1 and is_recycled = 0 and `order_id` > ? and is_top = 1"+
			" order by `order_id` asc limit 1", orderId).Scan(&article).Error
	} else {
		err = db.Db.Raw("select * from `articles`"+
			" where is_published = 1 and is_recycled = 0 and `order_id` > ? and is_top = 0"+
			" order by `order_id` asc limit 1", orderId).Scan(&article).Error
	}

	return
}

// GetByPage 分页获取文章
func (Article) GetByPage(page *utils.Pagination, key string, state uint,
	categoryId uint) ([]Article, uint, error) {
	var list []Article

	query := db.Db.Preload("User").Preload("Category").Preload("TagList").
		Model(&Article{}).Order("is_top desc,order_id asc,created_at desc", true)

	if key != "" {
		query = query.Where("title like concat('%',?,'%') or md_content like concat('%',?,'%')", key, key)
	}

	if state > 0 {
		switch state {
		// 已发布
		case 1:
			query = query.Where("is_published = 1 and is_recycled = 0")
		// 回收站
		case 2:
			query = query.Where("is_recycled = 1")
		// 草稿
		case 3:
			query = query.Where("is_published = 0 and is_recycled = 0")
		// 加密
		case 4:
			query = query.Where("is_published = 1 and is_recycled = 0 and pwd != ''")
		case 5:
			query = query.Where("is_published = 1 and is_recycled = 0")
		default:
			break
		}
	}

	if categoryId > 0 {
		query = query.Where("category_id = ?", categoryId)
	}
	query = query.Select([]string{"id", "created_at", "user_id", "category_id", "order_id", "is_top", "is_recycled",
		"is_published", "pwd", "url", "title", "summary", "img", "comment_count", "visit_count"})
	total, err := utils.ToPage(page, query, &list)

	return list, total, err
}

// GetByUrl 根据 Url 获取文章
func (Article) GetByUrl(url string) (article Article, err error) {
	err = db.Db.Preload("Category").Preload("TagList").Where("`url` = ?", url).First(&article).Error
	return
}

// Create 添加文章
func (article Article) Create(tagIds string) error {
	// markdown 渲染
	article.MDContent = setting.LuteEngine.MarkdownStr("", article.Content)

	// 若摘要为空，截取文章前 100 个字作为摘要
	if article.Summary == "" {
		content := []rune(utils.GetHtmlContent(article.MDContent))
		if len(content) < 100 {
			article.Summary = string(content)
		} else {
			article.Summary = string(content[:100])
		}
	}

	// 若图片为空，设置默认图片
	if article.Img == "" {
		article.Img = "https://s1.ax1x.com/2020/06/29/NWtFJA.jpg"
	}

	// 若 URL 为空，设置默认 URL
	if article.URL == "" {
		article.URL = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	}

	// 若为密码不为空，加密密码
	if article.Pwd != "" {
		var err error
		article.Pwd, err = utils.EncryptPwd(article.Pwd)
		if err != nil {
			return err
		}
	}

	var maxOrderId *uint
	err := db.Db.Raw("select MAX(`order_id`) `maxOrderId` from `articles`").
		Row().Scan(&maxOrderId)
	if err != nil {
		return err
	}

	if maxOrderId == nil {
		article.OrderId = 1
	} else {
		article.OrderId = *maxOrderId + 1
	}

	// 开始事务
	tx := db.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	article.CreatedAt = time.Now()

	// 添加文章
	err = tx.Exec("INSERT INTO `articles` (`user_id`, `category_id`, `order_id`, `is_top`, `is_recycled`,"+
		" `is_published`, `is_allow_commented`, `pwd`, `url`, `title`,"+
		" `summary`, `img`, `content`, `md_content`, `keywords`, `created_at`)"+
		" VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		article.UserId, article.CategoryId, article.OrderId, article.IsTop, article.IsRecycled,
		article.IsPublished, article.IsAllowCommented, article.Pwd, article.URL, article.Title,
		article.Summary, article.Img, article.Content, article.MDContent, article.Keywords, article.CreatedAt,
	).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 根据标题获取文章
	err = tx.Where("title = ?", article.Title).First(&article).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新分类对应文章数量
	err = tx.Exec("update `categories` set `count` = `count` + 1 where `id` = ?", article.CategoryId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 创建文章和标签关联，更新标签对应文章数量
	if tagIds != "" {
		tagIdList := strings.Split(tagIds, ",") // 根据 , 分割成字符串数组
		for _, tagId := range tagIdList {
			err = tx.Exec("insert into `tag_article` (`article_id`,`tag_id`) values (?,?)",
				article.ID, tagId).Error
			if err != nil {
				tx.Rollback()
				return err
			}

			err = tx.Exec("update `tags` set `count` = `count` + 1 where `id` = ?", tagId).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// Update 更新文章
func (article Article) Update(tagIds string) error {
	// markdown 渲染
	article.MDContent = setting.LuteEngine.MarkdownStr("", article.Content)

	// 若摘要为空，截取文章前 100 个字作为摘要
	if article.Summary == "" {
		content := []rune(utils.GetHtmlContent(article.MDContent))
		if len(content) < 100 {
			article.Summary = string(content)
		} else {
			article.Summary = string(content[:100])
		}
	}

	// 若图片为空，设置默认图片
	if article.Img == "" {
		article.Img = "https://s1.ax1x.com/2020/06/29/NWtFJA.jpg"
	}

	// 若 URL 为空，设置默认 URL
	if article.URL == "" {
		article.URL = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	}

	// 开始事务
	tx := db.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	a := Article{}
	err := tx.Where("`id` = ?", article.ID).First(&a).Error
	if err != nil {
		return err
	}

	// 密码发生变化，更新密码
	if a.Pwd != article.Pwd {
		// 密码不为空，加密密码
		if article.Pwd != "" {
			var err error

			article.Pwd, err = utils.EncryptPwd(article.Pwd)
			if err != nil {
				return err
			}
		}
	}

	// 使用 map 来更新，避免 gorm 默认不更新值为 nil, false, 0 的字段
	err = tx.Model(&Article{}).Where("`id` = ?", article.ID).
		Updates(map[string]interface{}{
			"category_id":        article.CategoryId,
			"order_id":           article.OrderId,
			"is_top":             article.IsTop,
			"is_recycled":        article.IsRecycled,
			"is_published":       article.IsPublished,
			"is_allow_commented": article.IsAllowCommented,
			"url":                article.URL,
			"title":              article.Title,
			"summary":            article.Summary,
			"img":                article.Img,
			"pwd":                article.Pwd,
			"content":            article.Content,
			"md_content":         article.MDContent,
			"keywords":           article.Keywords,
		}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 分类改变，更新分类对应文章数量
	if a.CategoryId != article.CategoryId {
		err = tx.Exec("update `categories` set `count` = `count` - 1 where (`id` = ? and `count` > 0)", a.CategoryId).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		err = tx.Exec("update `categories` set `count` = `count` + 1 where `id` = ?", article.CategoryId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 更新原标签对应文章数量
	err = tx.Exec("update `tags` set `count` = `count` - 1 where (`id` in "+
		"(select `tag_id` from `tag_article` where `article_id` = ?) and `count` > 0)", article.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除标签文章表中关联记录
	err = tx.Exec("delete from `tag_article` where `article_id` = ?", article.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 在标签文章表中添加关联记录
	if tagIds != "" {
		tagIdList := strings.Split(tagIds, ",")
		for _, tagId := range tagIdList {
			err = tx.Exec("insert into `tag_article` (`article_id`,`tag_id`) values (?,?)",
				article.ID, tagId).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}

		// 更新现标签对应文章数量
		err = tx.Exec("update `tags` set `count` = `count` + 1 where `id` in (?)", tagIdList).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// UpdateVisitCount 更新访问量
func (article Article) UpdateVisitCount() error {
	// 使用 UpdateColumn 函数来更新单个字段，避免 gorm 自动更新 updated_at 字段
	return db.Db.Model(&Article{}).Where("id = ?", article.ID).
		UpdateColumn("visit_count", article.VisitCount+1).Error
}

// RecycleOrRecover 将文章加入回收站或恢复
func (article Article) RecycleOrRecover(id string) (err error) {
	err = db.Db.Exec("update articles set is_recycled = !is_recycled where id = ?", id).Error

	return
}

// MoveUp 向上移动文章
func (article Article) MoveUp(currId, preId, currOrderId, preOrderId uint) error {
	// 执行事务
	return db.Db.Transaction(func(tx *gorm.DB) error {
		// 交换当前文章和上一篇文章的 OrderId
		err := db.Db.Model(&Article{}).Where("`id` = ?", currId).
			Update("order_id", preOrderId).Error
		if err != nil {
			return err
		}

		err = db.Db.Model(&Article{}).Where("`id` = ?", preId).
			Update("order_id", currOrderId).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// MoveDown 向下移动文章
func (article Article) MoveDown(currId, nextId, currOrderId, nextOrderId uint) error {
	// 执行事务
	return db.Db.Transaction(func(tx *gorm.DB) error {
		// 交换当前文章和下一篇文章的 OrderId
		err := db.Db.Model(&Article{}).Where("`id` = ?", currId).
			Update("order_id", nextOrderId).Error
		if err != nil {
			return err
		}

		err = db.Db.Model(&Article{}).Where("`id` = ?", nextId).
			Update("order_id", currOrderId).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// DeleteById 删除文章
func (Article) DeleteById(id string) error {
	// 开始事务
	tx := db.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	// 更新分类对应文章数量
	err := tx.Exec("update `categories` set `count` = `count` - 1 where (`id` in "+
		"(select `category_id` from `articles` where `id` = ?) and `count` > 0)", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新标签对应文章数量
	err = tx.Exec("update `tags` set `count` = `count` - 1 where (`id` in "+
		"(select `tag_id` from `tag_article` where `article_id` = ?) and `count` > 0)", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除标签文章表中的记录
	err = tx.Exec("delete from `tag_article` where `article_id` = ?", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除文章表中的记录
	err = tx.Where("`id` = ?", id).Unscoped().Delete(&Article{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// MultiDelByIds 批量删除文章
func (Article) MultiDelByIds(ids string) error {
	// 开始事务
	tx := db.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	idList := strings.Split(ids, ",")

	// 更新分类对应文章数量
	err := tx.Exec("update `categories` set `count` = `count` - 1 where (`id` in "+
		"(select `category_id` from `articles` where `id` in (?)) and `count` > 0)", idList).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新标签对应文章数量
	err = tx.Exec("update `tags` set `count` = `count` - 1 where (`id` in "+
		"(select `tag_id` from `tag_article` where `article_id` in (?)) and `count` > 0)", idList).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除标签文章表中的记录
	err = tx.Exec("delete from `tag_article` where `article_id` in (?)", idList).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除文章表中的记录
	err = tx.Where("`id` in (?)", idList).Unscoped().Delete(&Article{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// SaveFromFile 从文件导入文章
func (article Article) SaveFromFile() (err error) {
	user := User{}

	err = db.Db.First(&user).Error
	if err != nil {
		return
	}

	article.UserId = user.ID
	article.MDContent = setting.LuteEngine.MarkdownStr("", article.Content)

	err = article.Create("")

	return
}

// GetAll 归档
func (Archive) GetAll() (list []Archive, err error) {
	err = db.Db.Raw("select year(a.created_at) `year`, month(a.created_at) `month`, count(*) `count`" +
		" from articles a" +
		" group by `year`, `month`" +
		" order by `year` desc, `month` desc").Scan(&list).Error

	return
}
