package form

import (
	"aries/model"
)

// 用户信息表单
type UserInfoForm struct {
	Username  string `json:"username" binding:"required,min=3,max=30" label:"用户名"`
	Nickname  string `json:"nickname" label:"昵称"`
	Email     string `json:"email" binding:"required,max=30,email" label:"邮箱"`
	Signature string `json:"signature" label:"个性签名"`
}

// 绑定用户信息表单到用户实体
func (form UserInfoForm) BindToModel() model.User {
	return model.User{
		Username:  form.Username,
		Email:     form.Email,
		Nickname:  form.Nickname,
		Signature: form.Signature,
	}
}
