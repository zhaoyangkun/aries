package models

import (
	"aries/config/db"
	"github.com/jinzhu/gorm"
	"strings"
)

// 菜单
type Nav struct {
	gorm.Model
	ChildNavs   []*Nav `gorm:"ForeignKey:ParentNavId" json:"child_navs"`    // 子级菜单
	ParentNavId uint   `json:"parent_nav_id"`                               // 父级菜单 ID
	OrderId     uint   `gorm:"type:int;default:1;" json:"order_id"`         // 排序 ID
	OpenType    uint   `gorm:"type:tinyint(1);default:0;" json:"open_type"` // 打开方式，0 表示当前窗口，1 表示新窗口
	Name        string `gorm:"varchar(100);not null;" json:"name"`          // 名称
	Url         string `gorm:"varchar(255);not null;" json:"url"`           // 访问地址
	Icon        string `gorm:"varchar(255);" json:"icon"`                   // 图标
}

// 获取所有菜单（子菜单包含在父菜单中）
func (Nav) GetAll() ([]Nav, error) {
	var childNavs []Nav
	var parentNavs []Nav
	err := db.Db.Where("parent_nav_id > 0").Order("order_id desc", true).
		Find(&childNavs).Error
	if err != nil {
		return parentNavs, err
	}
	err = db.Db.Where("parent_nav_id = 0").Order("order_id desc", true).
		Find(&parentNavs).Error
	if err != nil {
		return parentNavs, err
	}
	for i := range parentNavs {
		for j := range childNavs {
			if childNavs[j].ParentNavId == parentNavs[i].ID {
				parentNavs[i].ChildNavs = append(parentNavs[i].ChildNavs, &childNavs[j])
			}
		}
	}
	return parentNavs, nil
}

// 添加菜单
func (n *Nav) Create() error {
	var maxOrderId *uint
	err := db.Db.Raw("select MAX(`order_id`) `maxOrderId` from `navs`").
		Row().Scan(&maxOrderId)
	if err != nil {
		return err
	}
	n.OrderId = *maxOrderId
	return db.Db.Create(&n).Error
}

// 修改菜单
func (n *Nav) Update() error {
	return db.Db.Model(&Nav{}).Updates(&n).Error
}

// 获取前一个菜单
func (n *Nav) GetPre() (Nav, error) {
	preNav := Nav{}
	err := db.Db.Where("order_id < ?", n.OrderId).Order("order_id DESC", true).
		First(&preNav).Error
	return preNav, err
}

// 获取后一个菜单
func (n *Nav) GetNext() (Nav, error) {
	nextNav := Nav{}
	err := db.Db.Where("order_id > ?", n.OrderId).Order("order_id ASC", true).
		First(&nextNav).Error
	return nextNav, err
}

// 向上移动菜单
func (n *Nav) MoveUp(currId, preId, currOrderId, preOrderId uint) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		err := db.Db.Where("`id` = ?", currId).Update("order_id", preOrderId).Error
		if err != nil {
			return err
		}
		err = db.Db.Where("`id` = ?", preId).Update("order_id", currOrderId).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// 向下移动菜单
func (n *Nav) MoveDown(currId, nextId, currOrderId, nextOrderId uint) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		err := db.Db.Where("`id` = ?", currId).Update("order_id", nextOrderId).Error
		if err != nil {
			return err
		}
		err = db.Db.Where("`id` = ?", nextId).Update("order_id", currOrderId).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// 删除菜单
func (Nav) DeleteById(id string) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		err := db.Db.Where("`id` = ?", id).Unscoped().Delete(&Nav{}).Error
		if err != nil {
			return err
		}
		err = db.Db.Where("`parent_nav_id` = ?", id).Unscoped().Delete(&Nav{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// 批量删除菜单
func (Nav) MultiDelByIds(ids string) error {
	idList := strings.Split(ids, ",")
	return db.Db.Transaction(func(tx *gorm.DB) error {
		err := db.Db.Where("`id` in (?)", idList).Unscoped().Delete(&Nav{}).Error
		if err != nil {
			return err
		}
		err = db.Db.Where("`parent_nav_id` in (?)", idList).Unscoped().Delete(&Nav{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}
