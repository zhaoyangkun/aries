package api

import (
	"aries/config/setting"
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ThemeHandler struct {
}

// GetAllThemes
// @Summary 获取所有主题
// @Tags 主题
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_themes [get]
func (t *ThemeHandler) GetAllThemes(ctx *gin.Context) {
	list, err := models.Theme{}.GetAll()
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
		Data: list,
	})
}

// GetThemeByName
// @Summary 根据主题名称获取主题
// @Tags 主题
// @version 1.0
// @Accept application/json
// @Param name path string true "主题名称"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/themes/{name} [get]
func (t *ThemeHandler) GetThemeByName(ctx *gin.Context) {
	themeName := ctx.Param("name")
	theme, err := models.Theme{}.GetByThemeName(themeName)
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
		Data: theme,
	})
}

// EnableTheme
// @Summary 启用主题
// @Tags 主题
// @version 1.0
// @Accept application/json
// @Param ThemeForm body forms.ThemeForm true "主题表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/themes [post]
func (t *ThemeHandler) EnableTheme(ctx *gin.Context) {
	form := forms.ThemeForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	theme := form.BindToModel()
	if err := theme.EnableTheme(); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	s := models.SysSetting{}
	sysSetting, _ := s.GetByName("网站设置")
	themeNameItem := models.SysSettingItem{
		SysId: sysSetting.ID,
		Key:   "theme_name",
		Val:   form.ThemeName,
	}
	itemList := []models.SysSettingItem{themeNameItem}
	err := models.SysSettingItem{}.MultiCreateOrUpdate(itemList)
	if err != nil {
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
		Msg:  "启用新主题成功",
		Data: nil,
	})
}
