package api

import (
	"aries/forms"
	"aries/models"
	"aries/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type NavHandler struct {
}

// @Summary 获取所有菜单
// @Tags 菜单
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/navs [get]
func (n *NavHandler) GetAllNavs(ctx *gin.Context) {
	list, err := models.Nav{}.GetAll()
	if err != nil {
		log.Error("error: ", err.Error())
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

// @Summary 添加菜单
// @Tags 菜单
// @version 1.0
// @Accept application/json
// @Param addForm body forms.NavAddForm true "添加菜单表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/navs [post]
func (n *NavHandler) CreateNav(ctx *gin.Context) {
	addForm := forms.NavAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}
	nav := addForm.BindToModel()
	if err := nav.Create(); err != nil {
		log.Error("error: ", err.Error())
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

// @Summary 修改菜单
// @Tags 菜单
// @version 1.0
// @Accept application/json
// @Param addForm body forms.NavEditForm true "修改菜单表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/navs [put]
func (n *NavHandler) UpdateNav(ctx *gin.Context) {
	editForm := forms.NavEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}
	nav := editForm.BindToModel()
	if err := nav.Update(); err != nil {
		log.Error("error: ", err.Error())
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

// @Summary 删除菜单
// @Tags 菜单
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/navs/{id} [delete]
func (n *NavHandler) DeleteNav(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.Nav{}.DeleteById(id)
	if err != nil {
		log.Error("error: ", err.Error())
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

// @Summary 批量删除菜单
// @Tags 菜单
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/navs [delete]
func (n *NavHandler) MultiDelNavs(ctx *gin.Context) {
	ids := ctx.Query("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}
	err := models.Nav{}.MultiDelByIds(ids)
	if err != nil {
		log.Error("error: ", err.Error())
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
