package api

import (
	"aries/forms"
	"aries/handlers"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NavHandler struct {
}

// GetAllNavs
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

// CreateNav
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
	oldNav, _ := nav.GetByName(nav.Name)
	if oldNav.Name != "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "该菜单名称已存在",
			Data: nil,
		})
		return
	}

	if err := nav.Create(); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "添加成功",
		Data: nil,
	})
}

// UpdateNav
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

	if editForm.ParentNavId == editForm.ID {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "不能将自身设置为父级菜单",
			Data: nil,
		})
		return
	}

	nav := editForm.BindToModel()
	if err := nav.Update(); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "修改成功",
		Data: nil,
	})
}

// MoveNavUp
// @Summary 向上移动菜单
// @Tags 菜单
// @version 1.0
// @Accept application/json
// @Param order_id path int true "order_id"
// @Param type path string true "type"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/navs/{type}/up/{order_id} [patch]
func (n *NavHandler) MoveNavUp(ctx *gin.Context) {
	navType := ctx.Param("type")
	orderIdStr := ctx.Param("order_id")

	orderId, err := strconv.Atoi(orderIdStr)
	if (navType != "parent" && navType != "child") || err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请求参数错误",
			Data: nil,
		})
		return
	}

	currNav, err := models.Nav{}.GetByOrderId(uint(orderId))
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	preNav, _ := currNav.GetPre(navType)
	if preNav.Name == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "无法再向上移动",
			Data: nil,
		})
		return
	}

	err = currNav.MoveUp(preNav)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "向上移动成功",
		Data: nil,
	})
}

// MoveNavDown
// @Summary 向移下动菜单
// @Tags 菜单
// @version 1.0
// @Accept application/json
// @Param order_id path int true "order_id"
// @Param type path string true "type"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/navs/{type}/down/{order_id} [patch]
func (n *NavHandler) MoveNavDown(ctx *gin.Context) {
	navType := ctx.Param("type")
	orderIdStr := ctx.Param("order_id")

	orderId, err := strconv.Atoi(orderIdStr)
	if (navType != "parent" && navType != "child") || err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请求参数错误",
			Data: nil,
		})
		return
	}

	currNav, err := models.Nav{}.GetByOrderId(uint(orderId))
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	nextNav, _ := currNav.GetNext(navType)
	if nextNav.Name == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "无法再向下移动",
			Data: nil,
		})
		return
	}

	err = currNav.MoveDown(nextNav)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "向下移动成功",
		Data: nil,
	})
}

// DeleteNav
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
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// MultiDelNavs
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
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}
