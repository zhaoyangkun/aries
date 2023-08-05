package api

import (
	"aries/config/setting"
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/jinzhu/gorm"
)

type SysSettingHandler struct {
}

// GetBlogVars
// @Summary 获取博客全局变量
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/blog_vars [get]
func (s *SysSettingHandler) GetBlogVars(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "查询成功",
		Data: setting.BlogVars,
	})
}

// GetSysSettingItem
// @Summary 获取设置条目
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param name query string false "设置名称"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/items [get]
func (s *SysSettingHandler) GetSysSettingItem(ctx *gin.Context) {
	name := ctx.Query("name")

	result, _ := models.SysSettingItem{}.GetBySysSettingName(name)
	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "查询成功",
		Data: result,
	})
}

// SaveSiteSetting
// @Summary 保存网站配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param settingForm body forms.SiteSettingForm true "网站配置表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/site [post]
func (s *SysSettingHandler) SaveSiteSetting(ctx *gin.Context) {
	settingForm := forms.SiteSettingForm{}
	if err := ctx.ShouldBindJSON(&settingForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(settingForm.SysId, 10, 0)
	sysSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  settingForm.TypeName,
	}
	if sysId == 0 {
		staticRootVal := settingForm.SiteUrl
		if setting.Config.Server.Mode == gin.ReleaseMode {
			staticRootVal = "https://gcore.jsdelivr.net/gh/zhaoyangkun/aries@latest"
		}
		staticRootItem := models.SysSettingItem{
			SysId: sysSetting.ID,
			Key:   "static_root",
			Val:   staticRootVal,
		}
		sysSetting.Items = append(sysSetting.Items, staticRootItem)
		if err := sysSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	t := reflect.TypeOf(settingForm)
	v := reflect.ValueOf(settingForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: sysSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	blogSetting, _ := models.SysSettingItem{}.GetBySysSettingName("网站设置")
	socialInfo, _ := models.SysSettingItem{}.GetBySysSettingName("社交信息")
	setting.BlogVars.InitBlogVars(blogSetting, socialInfo)

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SaveSMTPSetting
// @Summary 保存 SMTP 服务配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param settingForm body forms.EmailSettingForm true "SMTP 配置表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/smtp [post]
func (s *SysSettingHandler) SaveSMTPSetting(ctx *gin.Context) {
	settingForm := forms.EmailSettingForm{}
	if err := ctx.ShouldBindJSON(&settingForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(settingForm.SysId, 10, 0)
	sysSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  settingForm.TypeName,
	}
	if sysId == 0 {
		if err := sysSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	settingForm.SysId = strconv.Itoa(int(sysSetting.ID))
	t := reflect.TypeOf(settingForm)
	v := reflect.ValueOf(settingForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: sysSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SaveQubuSetting
// @Summary 保存去不图床表单配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param QubuForm body forms.QubuForm true "去不图床表单配置信息"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/pic_bed/qubu [post]
func (s *SysSettingHandler) SaveQubuSetting(ctx *gin.Context) {
	qubuForm := forms.QubuForm{}
	if err := ctx.ShouldBindJSON(&qubuForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(qubuForm.SysId, 10, 0)
	qubuSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  qubuForm.StorageType,
	}
	picBedSetting := models.SysSetting{
		Name: "图床设置",
	}

	picBedSettingItems, _ := models.SysSettingItem{}.GetBySysSettingName("图床设置")
	if len(picBedSettingItems) == 0 {
		if err := picBedSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	} else {
		sysId, _ := strconv.Atoi(picBedSettingItems["sys_id"])
		picBedSetting.ID = uint(sysId)
	}
	log.Logger.Sugar().Info("picBedSetting: ", picBedSetting)

	if sysId == 0 {
		if err := qubuSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	picBedForm := forms.PicBedSettingForm{
		SysId:       strconv.Itoa(int(picBedSetting.ID)),
		StorageType: "7bu",
	}

	t := reflect.TypeOf(picBedForm)
	v := reflect.ValueOf(picBedForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: picBedSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	qubuForm.SysId = strconv.Itoa(int(qubuSetting.ID))
	t = reflect.TypeOf(qubuForm)
	v = reflect.ValueOf(qubuForm)
	itemList = []models.SysSettingItem{}
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: qubuSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err = models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SaveSmmsSetting
// @Summary 保存 sm.ms 配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param smmsForm body forms.SmmsForm true "sm.ms 配置表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/pic_bed/smms [post]
func (s *SysSettingHandler) SaveSmmsSetting(ctx *gin.Context) {
	smmsForm := forms.SmmsForm{}
	if err := ctx.ShouldBindJSON(&smmsForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(smmsForm.SysId, 10, 0)
	smmsSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  smmsForm.StorageType,
	}
	picBedSetting := models.SysSetting{
		Name: "图床设置",
	}

	picBedSettingItems, _ := models.SysSettingItem{}.GetBySysSettingName("图床设置")
	if len(picBedSettingItems) == 0 {
		if err := picBedSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	} else {
		sysId, _ := strconv.Atoi(picBedSettingItems["sys_id"])
		picBedSetting.ID = uint(sysId)
	}
	log.Logger.Sugar().Info("picBedSetting: ", picBedSetting)

	if sysId == 0 {
		if err := smmsSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	picBedForm := forms.PicBedSettingForm{
		SysId:       strconv.Itoa(int(picBedSetting.ID)),
		StorageType: "sm.ms",
	}

	t := reflect.TypeOf(picBedForm)
	v := reflect.ValueOf(picBedForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: picBedSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	smmsForm.SysId = strconv.Itoa(int(smmsSetting.ID))
	t = reflect.TypeOf(smmsForm)
	v = reflect.ValueOf(smmsForm)
	itemList = []models.SysSettingItem{}
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: smmsSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err = models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SaveImgbbSetting
// @Summary 保存 imgbb 配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param imgbbForm body forms.ImgbbForm true "imgbb 配置表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/pic_bed/imgbb [post]
func (s *SysSettingHandler) SaveImgbbSetting(ctx *gin.Context) {
	imgbbForm := forms.ImgbbForm{}
	if err := ctx.ShouldBindJSON(&imgbbForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(imgbbForm.SysId, 10, 0)
	imgbbSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  imgbbForm.StorageType,
	}
	picBedSetting := models.SysSetting{
		Name: "图床设置",
	}

	picBedSettingItems, _ := models.SysSettingItem{}.GetBySysSettingName("图床设置")
	if len(picBedSettingItems) == 0 {
		if err := picBedSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	} else {
		sysId, _ := strconv.Atoi(picBedSettingItems["sys_id"])
		picBedSetting.ID = uint(sysId)
	}

	if sysId == 0 {
		if err := imgbbSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}
	picBedForm := forms.PicBedSettingForm{
		SysId:       strconv.Itoa(int(picBedSetting.ID)),
		StorageType: "imgbb",
	}
	t := reflect.TypeOf(picBedForm)
	v := reflect.ValueOf(picBedForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: picBedSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	imgbbForm.SysId = strconv.Itoa(int(imgbbSetting.ID))
	t = reflect.TypeOf(imgbbForm)
	v = reflect.ValueOf(imgbbForm)
	itemList = []models.SysSettingItem{}
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: imgbbSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err = models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SaveTencentCosSetting
// @Summary 保存腾讯云 COS 配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param cosForm body forms.TencentCosForm true "腾讯云 COS 配置表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/pic_bed/tencent_cos [post]
func (s *SysSettingHandler) SaveTencentCosSetting(ctx *gin.Context) {
	cosForm := forms.TencentCosForm{}
	if err := ctx.ShouldBindJSON(&cosForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(cosForm.SysId, 10, 0)
	cosSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  cosForm.StorageType,
	}
	picBedSetting := models.SysSetting{
		Name: "图床设置",
	}

	picBedSettingItems, _ := models.SysSettingItem{}.GetBySysSettingName("图床设置")
	if len(picBedSettingItems) == 0 {
		if err := picBedSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	} else {
		sysId, _ := strconv.Atoi(picBedSettingItems["sys_id"])
		picBedSetting.ID = uint(sysId)
	}

	if sysId == 0 {
		if err := cosSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	picBedForm := forms.PicBedSettingForm{
		SysId:       strconv.Itoa(int(picBedSetting.ID)),
		StorageType: "cos",
	}
	t := reflect.TypeOf(picBedForm)
	v := reflect.ValueOf(picBedForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: picBedSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	cosForm.SysId = strconv.Itoa(int(cosSetting.ID))
	t = reflect.TypeOf(cosForm)
	v = reflect.ValueOf(cosForm)
	itemList = []models.SysSettingItem{}
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: cosSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err = models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SendTestEmail
// @Summary 发送测试邮件
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// Param sendForm body orm.EmailSendForm true "发送邮件表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/email/test [post]
func (s *SysSettingHandler) SendTestEmail(ctx *gin.Context) {
	sendForm := forms.EmailSendForm{}
	if err := ctx.ShouldBindJSON(&sendForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	emailSetting, err := models.SysSettingItem{}.GetBySysSettingName("邮件设置")
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	siteItems, err := models.SysSettingItem{}.GetBySysSettingName("网站设置")
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	if len(emailSetting) == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
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
	msg.SetBody("text/html", utils.GetEmailHTML(sendForm.Title, siteItems["site_url"], sendForm.ReceiveEmail,
		sendForm.Content))
	port, _ := strconv.Atoi(emailSetting["port"])
	// 设置 SMTP 参数
	d := gomail.NewDialer(emailSetting["address"], port, emailSetting["account"], emailSetting["pwd"])

	// 发送
	err = d.DialAndSend(msg)
	if err != nil {
		log.Logger.Sugar().Error("邮件发送失败：", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "邮件发送失败",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "发送成功",
		Data: nil,
	})
}

// GetAdminIndexData
// @Summary 获取后台首页数据
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/index_info [get]
func (s *SysSettingHandler) GetAdminIndexData(ctx *gin.Context) {
	articleCount, err := models.Article{}.GetCount()
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	commentCount, err := models.Comment{}.GetCount()
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	latestArticles, err := models.Article{}.GetLatest(6)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	latestComments, err := models.Comment{}.GetLatest(6)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "查询成功",
		Data: gin.H{
			"article_count":   articleCount,
			"comment_count":   commentCount,
			"latest_articles": latestArticles,
			"latest_comments": latestComments,
		},
	})
}

// SaveLocalCommentSetting
// @Summary 保存评论配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param settingForm body forms.CommentSettingForm true "评论配置表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /sys_setting/comment/local [post]
func (s *SysSettingHandler) SaveLocalCommentSetting(ctx *gin.Context) {
	localCommentForm := forms.LocalCommentSettingForm{}
	if err := ctx.ShouldBindJSON(&localCommentForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(localCommentForm.SysId, 10, 0)
	localCommentSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  localCommentForm.PlugIn,
	}
	commentPlugInSetting := models.SysSetting{
		Name: "评论组件设置",
	}

	commentPlugInSettingItems, _ := models.SysSettingItem{}.GetBySysSettingName("评论组件设置")
	if len(commentPlugInSettingItems) == 0 {
		if err := commentPlugInSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	} else {
		sysId, _ := strconv.Atoi(commentPlugInSettingItems["sys_id"])
		commentPlugInSetting.ID = uint(sysId)
	}
	log.Logger.Sugar().Info("commentPlugInSetting: ", commentPlugInSetting)

	if sysId == 0 {
		if err := localCommentSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	CommentPlugInForm := forms.CommentPlugInForm{
		SysId:  strconv.Itoa(int(commentPlugInSetting.ID)),
		PlugIn: "local-comment",
	}

	t := reflect.TypeOf(CommentPlugInForm)
	v := reflect.ValueOf(CommentPlugInForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: commentPlugInSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	localCommentForm.SysId = strconv.Itoa(int(localCommentSetting.ID))
	t = reflect.TypeOf(localCommentForm)
	v = reflect.ValueOf(localCommentForm)
	itemList = []models.SysSettingItem{}
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: localCommentSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err = models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SaveTwikooSetting
// @Summary 保存 Twikoo 配置信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param settingForm body forms.TwikooForm true "Twikoo 配置表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/comment/twikoo [post]
func (s *SysSettingHandler) SaveTwikooSetting(ctx *gin.Context) {
	twikooSettingForm := forms.TwikooSettingForm{}
	if err := ctx.ShouldBindJSON(&twikooSettingForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(twikooSettingForm.SysId, 10, 0)
	commentPlugInSetting := models.SysSetting{
		Name: "评论组件设置",
	}
	twikooSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  twikooSettingForm.PlugIn,
	}

	commentPlugInSettingItems, _ := models.SysSettingItem{}.GetBySysSettingName("评论组件设置")
	if len(commentPlugInSettingItems) == 0 { // 若不存在评论组件设置数据，则创建
		if err := commentPlugInSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	} else {
		sysId, _ := strconv.Atoi(commentPlugInSettingItems["sys_id"])
		commentPlugInSetting.ID = uint(sysId)
	}

	if sysId == 0 { // 不存在 twikoo 设置数据，则创建
		if err := twikooSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	// 保存评论组件相关设置
	commentPlugInForm := forms.CommentPlugInForm{
		SysId:  strconv.Itoa(int(commentPlugInSetting.ID)),
		PlugIn: "twikoo-comment",
	}
	t := reflect.TypeOf(commentPlugInForm)
	v := reflect.ValueOf(commentPlugInForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: commentPlugInSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	// 保存 twikoo 评论组件相关设置
	t = reflect.TypeOf(twikooSettingForm)
	v = reflect.ValueOf(twikooSettingForm)
	itemList = []models.SysSettingItem{}
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: twikooSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err = models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SaveParamSetting
// @Summary 保存参数配置
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param settingForm body forms.ParamSettingForm true "参数配置表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/param [post]
func (s *SysSettingHandler) SaveParamSetting(ctx *gin.Context) {
	settingForm := forms.ParamSettingForm{}
	if err := ctx.ShouldBindJSON(&settingForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	sysId, _ := strconv.ParseUint(settingForm.SysId, 10, 0)
	sysSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  settingForm.TypeName,
	}

	if sysId == 0 {
		if err := sysSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}
	t := reflect.TypeOf(settingForm)
	v := reflect.ValueOf(settingForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: sysSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// SaveSocialInfo
// @Summary 保存社交信息
// @Tags 系统设置
// @version 1.0
// @Accept application/json
// @Param settingForm body forms.SocialInfoForm true "社交信息表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/sys_setting/social_info [post]
func (s *SysSettingHandler) SaveSocialInfo(ctx *gin.Context) {
	settingForm := forms.SocialInfoForm{}
	_ = ctx.ShouldBindJSON(&settingForm)

	sysId, _ := strconv.ParseUint(settingForm.SysId, 10, 0)
	sysSetting := models.SysSetting{
		Model: gorm.Model{ID: uint(sysId)},
		Name:  settingForm.TypeName,
	}

	if sysId == 0 {
		if err := sysSetting.Create(); err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}
	t := reflect.TypeOf(settingForm)
	v := reflect.ValueOf(settingForm)
	var itemList []models.SysSettingItem
	for i := 0; i < t.NumField(); i++ {
		item := models.SysSettingItem{
			SysId: sysSetting.ID,
			Key:   t.Field(i).Tag.Get("json"),
			Val:   v.Field(i).Interface().(string),
		}
		itemList = append(itemList, item)
	}

	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	blogSetting, _ := models.SysSettingItem{}.GetBySysSettingName("网站设置")
	socialInfo, _ := models.SysSettingItem{}.GetBySysSettingName("社交信息")
	setting.BlogVars.InitBlogVars(blogSetting, socialInfo)

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}
