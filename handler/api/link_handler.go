package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// @Summary 获取所有友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/all_links [get]
func GetAllLinks(ctx *gin.Context) {
	list, err := model.Link{}.GetAll()
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "查询成功",
		Data: list,
	})
}

// @Summary 分页获取友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// Param pageForm query orm.LinkPageForm true "友链分页表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/links [get]
func GetLinksByPage(ctx *gin.Context) {
	pageForm := form.LinkPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)
	list, total, err := model.Link{}.GetByPage(&pageForm.Pagination, pageForm.Key, pageForm.CategoryId)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "查询成功",
		Data: util.GetPageData(list, total, pageForm.Pagination),
	})
}

// @Summary 添加友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// Param addForm body orm.LinkPageForm true "友链分页表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/links [post]
func CreateLink(ctx *gin.Context) {
	addForm := form.LinkAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	link := addForm.BindToModel()
	if err := link.Create(); err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "添加成功",
		Data: nil,
	})
}

// @Summary 修改友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// Param addForm body orm.LinkPageForm true "友链分页表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/links [post]
func UpdateLink(ctx *gin.Context) {
	editForm := form.LinkEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	link := editForm.BindToModel()
	if err := link.Update(); err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "修改成功",
		Data: nil,
	})
}

// @Summary 删除友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// Param id path int true "ID"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/links/{id} [delete]
func DeleteLink(ctx *gin.Context) {
	id := ctx.Param("id")
	err := model.Link{}.DeleteById(id)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// @Summary 删除友链
// @Tags 友链
// @version 1.0
// @Accept application/json
// @Param ids query string true "IDS"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/links [delete]
func MultiDelLinks(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "")
	if ids == "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请求参数错误",
			Data: nil,
		})
		return
	}
	err := model.Link{}.MultiDelByIds(ids)
	if err != nil {
		log.Errorln("Error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}
