package model

import "github.com/jinzhu/gorm"

// 导航栏结构
type Nav struct {
	gorm.Model
	ParentNav   *Nav   `gorm:"foreignkey:ParentNavId" json:"parent_nav"`    // 父级菜单
	ParentNavId uint   `json:"parent_nav_id"`                               // 父级菜单 ID
	OrderId     uint   `gorm:"type:int;default:1;" json:"order_id"`         // 排序 ID
	OpenType    uint   `gorm:"type:tinyint(1);default:0;" json:"open_type"` // 打开方式，0 表示当前窗口，1 表示新窗口
	Name        string `gorm:"varchar(100);not null;" json:"name"`          // 名称
	Url         string `gorm:"varchar(255);not null;" json:"url"`           // 访问地址
	Icon        string `gorm:"varchar(255);not null;" json:"icon"`          // 图标
}
