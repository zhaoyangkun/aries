package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
)

// 保存 SMTP 服务配置信息
func SaveSMTP(ctx *gin.Context) {
	emailForm := form.EmailForm{}
	if err := ctx.ShouldBindJSON(&emailForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	sysSetting := model.SysSetting{Name: emailForm.TypeName}
	err := sysSetting.Create()
	if err != nil {
		log.Errorln("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
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
	err = model.SysSettingItem{}.MultiCreate(itemList)
	if err != nil {
		log.Errorln("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "保存成功",
		Data: nil,
	})
}

// 发送邮件
func SendEmail(ctx *gin.Context) {

}
