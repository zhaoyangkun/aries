package forms

import "aries/utils"

type PicturePageForm struct {
	Key              string `form:"key"`          // 关键词
	StorageName      string `form:"storage_name"` //存储类型
	utils.Pagination        // 分页结构
}
