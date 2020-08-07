package models

import (
	"aries/config/db"
	"github.com/jinzhu/gorm"
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
	Icon        string `gorm:"varchar(255);not null;" json:"icon"`          // 图标
}

// 获取所有菜单
func (Nav) GetAll() ([]Nav, error) {
	var childNavs []Nav
	var allNavs []Nav
	err := db.Db.Where("parent_nav_id > 0").Find(&childNavs).Error
	if err != nil {
		return allNavs, err
	}
	err = db.Db.Find(&allNavs).Error
	if err != nil {
		return allNavs, err
	}
	for i := range allNavs {
		if allNavs[i].ParentNavId == 0 {
			for j := range childNavs {
				if childNavs[j].ParentNavId == allNavs[i].ID {
					allNavs[i].ChildNavs = append(allNavs[i].ChildNavs, &childNavs[j])
				}
			}
		}
	}
	return allNavs, nil
}
