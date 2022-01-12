package forms

import (
	"aries/models"
)

type ThemeForm struct {
	ThemeName string `json:"theme_name" binding:"required,max=100" label:"主题名称"`
}

func (form ThemeForm) BindToModel() models.Theme {
	return models.Theme{
		ThemeInfo: models.ThemeInfo{
			ThemeName: form.ThemeName,
		},
	}
}
