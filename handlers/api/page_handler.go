package api

import (
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageHandler struct {
}

// GetAllPages 获取所有页面
func (p *PageHandler) GetAllPages(ctx *gin.Context) {
	list, err := models.Page{}.GetAll()
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

// GetPagesByPage 分页获取页面
func (p *PageHandler) GetPagesByPage(ctx *gin.Context) {
	pageForm := forms.PageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)

	list, total, err := models.Page{}.GetByPage(&pageForm.Pagination, pageForm.Key)
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
		Data: utils.GetPageData(list, total, pageForm.Pagination),
	})
}

// CreatePage 创建页面
func (p *PageHandler) CreatePage(ctx *gin.Context) {
	addForm := forms.AddPageForm{}

	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	page := addForm.BindToModel()

	if err := page.Create(); err != nil {
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
		Msg:  "创建成功",
		Data: nil,
	})
}

// UpdatePage 修改页面
func (p *PageHandler) UpdatePage(ctx *gin.Context) {
	editForm := forms.EditPageForm{}

	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	page := editForm.BindToModel()

	if err := page.Update(); err != nil {
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
		Msg:  "修改成功",
		Data: nil,
	})
}

// MultiDelPages 批量删除页面
func (p *PageHandler) MultiDelPages(ctx *gin.Context) {
	ids := ctx.Query("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请选择要删除的页面",
			Data: nil,
		})
		return
	}

	err := models.Page{}.MultiDelByIds(ids)
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
		Msg:  "删除成功",
		Data: nil,
	})
}
