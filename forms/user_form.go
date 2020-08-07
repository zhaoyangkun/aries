package forms

import (
	"aries/models"
)

// 用户信息表单
type UserInfoForm struct {
	Username  string `json:"username" binding:"required,min=3,max=30" label:"用户名"`
	Nickname  string `json:"nickname" label:"昵称"`
	Email     string `json:"email" binding:"required,max=30,email" label:"邮箱"`
	Signature string `json:"signature" label:"个性签名"`
}

// 修改密码表单
type PwdForm struct {
	Username      string `json:"username" binding:"required,min=3,max=30" label:"用户名"`
	OldPwd        string `json:"old_pwd" binding:"required,min=6,max=20" label:"旧密码"`
	NewPwd        string `json:"new_pwd" binding:"required,min=6,max=20,nefield=OldPwd" label:"新密码"`
	ConfirmNewPwd string `json:"confirm_new_pwd" binding:"required,min=6,max=20,eqfield=NewPwd" label:"确认密码"`
}

// 绑定用户信息表单到用户实体
func (form UserInfoForm) BindToModel() models.User {
	return models.User{
		Username:  form.Username,
		Email:     form.Email,
		Nickname:  form.Nickname,
		Signature: form.Signature,
	}
}
