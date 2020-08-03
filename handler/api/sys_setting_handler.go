package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"strconv"
)

// @Summary 获取设置条目
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param name query string false "设置名称"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/sys_setting_items [get]
func GetSysSettingItem(ctx *gin.Context) {
	name := ctx.Query("name")
	result, _ := model.SysSettingItem{}.GetBySysSettingName(name)
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "查询成功",
		Data: result,
	})
}

// @Summary 保存网站配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param siteForm body form.SiteForm true "网站配置表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/site_setting [post]
func SaveSiteSetting(ctx *gin.Context) {
	siteForm := form.SiteForm{}
	if err := ctx.ShouldBindJSON(&siteForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	sysId, _ := strconv.ParseUint(siteForm.SysId, 10, 0)
	sysSetting := model.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  siteForm.TypeName,
	}
	if sysId == 0 {
		if err := sysSetting.Create(); err != nil {
			log.Errorln("error: ", err.Error())
			ctx.JSON(http.StatusOK, util.Result{
				Code: util.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}
	t := reflect.TypeOf(siteForm)
	v := reflect.ValueOf(siteForm)
	var itemList []model.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := model.SysSettingItem{
			SysId: sysSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}
	err := model.SysSettingItem{}.MultiCreateOrUpdate(sysSetting.ID, itemList)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// @Summary 保存 SMTP 服务配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param emailForm body form.EmailForm true "SMTP 配置表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/smtp_setting [post]
func SaveSMTPSetting(ctx *gin.Context) {
	emailForm := form.EmailForm{}
	if err := ctx.ShouldBindJSON(&emailForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	sysId, _ := strconv.ParseUint(emailForm.SysId, 10, 0)
	sysSetting := model.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  emailForm.TypeName,
	}
	if sysId == 0 {
		if err := sysSetting.Create(); err != nil {
			log.Errorln("error: ", err.Error())
			ctx.JSON(http.StatusOK, util.Result{
				Code: util.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}
	emailForm.SysId = strconv.Itoa(int(sysSetting.ID))
	t := reflect.TypeOf(emailForm)
	v := reflect.ValueOf(emailForm)
	var itemList []model.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := model.SysSettingItem{
			SysId: sysSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}
	err := model.SysSettingItem{}.MultiCreateOrUpdate(sysSetting.ID, itemList)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// @Summary 发送测试邮件
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// Param sendForm body orm.EmailSendForm true "发送邮件表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/test_send_email [post]
func SendTestEmail(ctx *gin.Context) {
	sendForm := form.EmailSendForm{}
	if err := ctx.ShouldBindJSON(&sendForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	emailSetting, err := model.SysSettingItem{}.GetBySysSettingName("邮件设置")
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	if len(emailSetting) == 0 {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请先配置 SMTP，再进行邮件测试",
			Data: nil,
		})
		return
	}
	msg := gomail.NewMessage()
	// 设置收件人
	msg.SetHeader("To", sendForm.ReceiveEmail)
	// 设置发件人
	msg.SetAddressHeader("From", emailSetting["account"], sendForm.Sender)
	// 主题
	msg.SetHeader("Subject", sendForm.Title)
	// 正文
	msg.SetBody("text/html", util.GetEmailHTML(sendForm.Title, sendForm.ReceiveEmail,
		sendForm.Content))
	port, _ := strconv.Atoi(emailSetting["port"])
	// 设置 SMTP 参数
	d := gomail.NewDialer(emailSetting["address"], port, emailSetting["account"], emailSetting["pwd"])
	// 发送
	err = d.DialAndSend(msg)
	if err != nil {
		log.Error("邮件发送失败：", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "邮件发送失败",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "发送成功",
		Data: nil,
	})
}

// 获取后台首页所需数据
func GetAdminIndexData(ctx *gin.Context) {
	articleCount, err := model.Article{}.GetCount()
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	commentCount, err := model.Comment{}.GetCount()
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	latestArticles, err := model.Article{}.GetLatest(10)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	latestComments, err := model.Comment{}.GetLatest(10)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "查询成功",
		Data: gin.H{
			"article_count":   articleCount,
			"comment_count":   commentCount,
			"latest_articles": latestArticles,
			"latest_comments": latestComments,
		},
	})
}
