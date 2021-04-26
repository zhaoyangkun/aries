package api

import (
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
}

// GetAllLinks
// @Summary 获取所有友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_links [get]
func (l *LinkHandler) GetAllLinks(ctx *gin.Context) {
	list, err := models.Link{}.GetAll()
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

// GetLinksByPage
// @Summary 分页获取友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// Param pageForm query orm.LinkPageForm true "友链分页表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/links [get]
func (l *LinkHandler) GetLinksByPage(ctx *gin.Context) {
	pageForm := forms.LinkPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)

	list, total, err := models.Link{}.GetByPage(&pageForm.Pagination, pageForm.Key, pageForm.CategoryId)
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

// CreateLink
// @Summary 添加友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// Param addForm body orm.LinkPageForm true "友链分页表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/links [post]
func (l *LinkHandler) CreateLink(ctx *gin.Context) {
	addForm := forms.LinkAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	link := addForm.BindToModel()
	if err := link.Create(); err != nil {
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
		Msg:  "添加成功",
		Data: nil,
	})
}

// UpdateLink
// @Summary 修改友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// Param addForm body orm.LinkPageForm true "友链分页表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/links [post]
func (l *LinkHandler) UpdateLink(ctx *gin.Context) {
	editForm := forms.LinkEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	link := editForm.BindToModel()
	if err := link.Update(); err != nil {
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

// DeleteLink
// @Summary 删除友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// Param id path int true "ID"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/links/{id} [delete]
func (l *LinkHandler) DeleteLink(ctx *gin.Context) {
	id := ctx.Param("id")

	err := models.Link{}.DeleteById(id)
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

// MultiDelLinks
// @Summary 删除友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// @Param ids query string true "IDS"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/links [delete]
func (l *LinkHandler) MultiDelLinks(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "")

	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请求参数错误",
			Data: nil,
		})
		return
	}

	err := models.Link{}.MultiDelByIds(ids)
	if err != nil {
		log.Logger.Sugar().Error("Error: ", err.Error())
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
