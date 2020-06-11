package model

import (
	"aries/config/db"
	"aries/util"
	"github.com/jinzhu/gorm"
)

// 用户结构体
type User struct {
	gorm.Model
	Username         string `gorm:"type:varchar(30);not null;" json:"username"`        // 用户名
	Pwd              string `gorm:"type:varchar(100);not null;" json:"pwd"`            // 密码
	CommentCheckedOn bool   `gorm:"type:bool;default:true;" json:"comment_checked_on"` // 开启评论审核：默认为 true
	UserImg          string `gorm:"type:varchar(255);not null;" json:"user_img"`       // 用户头像
	Nickname         string `gorm:"type:varchar(30);" json:"nickname"`                 // 昵称
	Email            string `gorm:"type:varchar(30);not null;" json:"email"`           // 邮箱
	SiteUrl          string `gorm:"type:varchar(255);" json:"site_url"`                // 网址
	Signature        string `gorm:"type:varchar(255);" json:"signature"`               // 个性签名
	QQ               string `gorm:"type:varchar(30);" json:"qq"`                       // qq 号
	WeChat           string `gorm:"type:varchar(30);" json:"we_chat"`                  // 微信号
	Github           string `gorm:"type:varchar(30);" json:"github"`                   // github 账号
	WeiBo            string `gorm:"type:varchar(30);" json:"wei_bo"`                   // 微博账号
	WeChatPayImg     string `gorm:"type:varchar(255);" json:"we_chat_pay_img"`         // 微信支付收款二维码
	AliPayImg        string `gorm:"type:varchar(255);" json:"ali_pay_img"`             // 支付宝收款二维码
}

// 初始化数据表
func init() {
	db.Db.AutoMigrate(User{})
}

// 获取所有用户
func (user User) GetAll() ([]User, error) {
	var users []User
	err := db.Db.Find(&users).Error
	return users, err
}

// 根据用户名和密码获取用户
func (user User) GetByUsername() (User, error) {
	var u User
	err := db.Db.Where("username = ?", user.Username).First(&u).Error
	return u, err
}

// 创建用户
func (user User) Create() error {
	hashedPwd, err := util.EncryptPwd(user.Pwd) // 加密密码
	if err != nil {
		return err
	}
	user.Pwd = hashedPwd
	return db.Db.Create(&user).Error
}
