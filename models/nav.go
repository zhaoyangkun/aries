package models

import (
	"aries/config/db"
	"strings"

	"github.com/jinzhu/gorm"
)

// Nav 菜单
type Nav struct {
	gorm.Model
	ChildNavs   []*Nav `gorm:"ForeignKey:ParentNavId" json:"child_navs"`    // 子级菜单
	ParentNavId uint   `json:"parent_nav_id"`                               // 父级菜单 ID
	OrderId     uint   `gorm:"type:int;default:1;" json:"order_id"`         // 排序 ID
	OpenType    uint   `gorm:"type:tinyint(1);default:0;" json:"open_type"` // 打开方式，0 表示当前窗口，1 表示新窗口
	Name        string `gorm:"type:varchar(100);not null;" json:"name"`     // 名称
	Url         string `gorm:"type:varchar(255);not null;" json:"url"`      // 访问地址
	Icon        string `gorm:"type:varchar(255);" json:"icon"`              // 图标
}

// GetByOrderId 根据 OrderId 获取菜单
func (Nav) GetByOrderId(orderId uint) (nav Nav, err error) {
	err = db.Db.Where("`order_id` = ?", orderId).First(&nav).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetByName 根据名称获取菜单
func (Nav) GetByName(name string) (nav Nav, err error) {
	err = db.Db.Where("`name` = ?", name).First(&nav).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return
}

// GetAll 获取所有菜单（子菜单包含在父菜单中）
func (Nav) GetAll() ([]Nav, error) {
	var childNavs []Nav
	var parentNavs []Nav

	err := db.Db.Where("parent_nav_id > 0").Order("order_id ASC", true).
		Find(&childNavs).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return parentNavs, err
	}

	err = db.Db.Where("parent_nav_id = 0").Order("order_id ASC", true).
		Find(&parentNavs).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
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

// Create 添加菜单
func (n *Nav) Create() error {
	var maxOrderId *uint
	err := db.Db.Raw("select MAX(`order_id`) `maxOrderId` from `navs`").
		Row().Scan(&maxOrderId)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	if maxOrderId == nil {
		n.OrderId = 1
	} else {
		n.OrderId = *maxOrderId + 1
	}

	return db.Db.Create(&n).Error
}

// Update 修改菜单
func (n *Nav) Update() error {
	return db.Db.Model(&Nav{}).Updates(&n).Error
}

// GetPre 获取前一个菜单
func (n *Nav) GetPre(navType string) (Nav, error) {
	preNav := Nav{}
	var err error

	if navType == "parent" {
		err = db.Db.Where("`parent_nav_id` = 0 and `order_id` < ?", n.OrderId).
			Order("order_id DESC", true).First(&preNav).Error
	} else {
		err = db.Db.Where("`parent_nav_id` > 0 and `order_id` < ?", n.OrderId).
			Order("order_id DESC", true).First(&preNav).Error
	}
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return preNav, err
}

// GetNext 获取后一个菜单
func (n *Nav) GetNext(navType string) (Nav, error) {
	nextNav := Nav{}
	var err error

	if navType == "parent" {
		err = db.Db.Where("`parent_nav_id` = 0 and `order_id` > ?", n.OrderId).
			Order("order_id ASC", true).First(&nextNav).Error
	} else {
		err = db.Db.Where("`parent_nav_id` > 0 and `order_id` > ?", n.OrderId).
			Order("order_id ASC", true).First(&nextNav).Error
	}
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return nextNav, err
}

// MoveUp 向上移动菜单
func (n *Nav) MoveUp(preNav Nav) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		err := db.Db.Model(&Nav{}).Where("`id` = ?", n.ID).Update("order_id", preNav.OrderId).Error
		if err != nil {
			return err
		}

		err = db.Db.Model(&Nav{}).Where("`id` = ?", preNav.ID).Update("order_id", n.OrderId).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// MoveDown 向下移动菜单
func (n *Nav) MoveDown(nextNav Nav) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		err := db.Db.Model(&Nav{}).Where("`id` = ?", n.ID).Update("order_id", nextNav.OrderId).Error
		if err != nil {
			return err
		}

		err = db.Db.Model(&Nav{}).Where("`id` = ?", nextNav.ID).Update("order_id", n.OrderId).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// DeleteById 删除菜单
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

// MultiDelByIds 批量删除菜单
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
