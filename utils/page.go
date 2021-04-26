package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 分页默认参数
var (
	defaultPageSize    uint = 10 // 默认每页条数
	defaultCurrentPage uint = 1  // 默认页码
)

// Pagination 分页结构体
type Pagination struct {
	Ok         bool        // 是否出错
	Size       uint        `form:"size"` // 每页条数
	Page       uint        `form:"page"` // 页码
	Data       interface{} // 分页数据
	Total      uint        // 总数
	TotalPages uint        // 总页数
}

// ToPage 进行分页
func ToPage(p *Pagination, db *gorm.DB, list interface{}) (uint, error) {
	// 设置默认参数
	if p.Size < 1 {
		p.Size = defaultPageSize
	}
	if p.Page < 1 {
		p.Page = defaultCurrentPage
	}

	// 获取全部搜索数量
	var total uint
	err := db.Count(&total).Error
	if err != nil {
		return 0, err
	}
	offset := p.Size * (p.Page - 1)

	// 获取偏移量的数据
	err = db.Limit(p.Size).Offset(offset).Find(list).Error
	if err != nil {
		return 0, err
	}

	return total, err
}

// GetPageData 封装分页数据
func GetPageData(list interface{}, total uint, page Pagination) gin.H {
	// 计算总页数
	if total%page.Size == 0 {
		page.TotalPages = total / page.Size
	} else {
		page.TotalPages = (total / page.Size) + 1
	}
	// 返回 map
	return gin.H{
		"ok":          true,            // 是否成功
		"data":        list,            // 分页数据
		"total_num":   total,           // 总条数
		"total_pages": page.TotalPages, // 总页数
		"page":        page.Page,       // 页码
		"size":        page.Size,       // 每页条数
	}
}
