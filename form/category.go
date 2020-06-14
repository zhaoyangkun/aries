package form

import "aries/util"

type CategoryPageForm struct {
	Key string `form:"key"`
	util.Pagination
}
