package api

import (
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JournalHandler struct {
}

// GetAllJournals
// @Summary 获取所有日志
// @Tags 日志
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_journals [get]
func (j *JournalHandler) GetAllJournals(ctx *gin.Context) {
	list, err := models.Journal{}.GetAll()
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

// GetJournalById
// @Summary 根据 ID 获取日志
// @Tags 日志
// @version 1.0
// @Accept application/json
// @Param id path uint true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/journals/{id} [get]
func (j *JournalHandler) GetJournalById(ctx *gin.Context) {
	id := ctx.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil || intId == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "id 必须为正整数",
			Data: nil,
		})
		return
	}

	obj, err := models.Journal{}.GetById(uint(intId))
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
		Data: obj,
	})
}

// GetJournalsByPage
// @Summary 分页获取日志
// @Tags 日志
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/journals [get]
func (j *JournalHandler) GetJournalsByPage(ctx *gin.Context) {
	pageForm := forms.JournalPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)

	list, total, err := models.Journal{}.GetByPage(&pageForm.Pagination, pageForm.Key)
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

// CreateJournal
// @Summary 创建日志
// @Tags 日志
// @version 1.0
// @Accept application/json
// @Param addForm body forms.JournalAddForm true "创建日志表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/journals [post]
func (j *JournalHandler) CreateJournal(ctx *gin.Context) {
	addForm := forms.JournalAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	journal := addForm.BindToModel()
	if err := journal.Create(); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "发表成功",
		Data: nil,
	})
}

// UpdateJournal
// @Summary 更新日志
// @Tags 日志
// @version 1.0
// @Accept application/json
// @Param editForm body forms.JournalEditForm true "更新日志表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/journals [put]
func (j *JournalHandler) UpdateJournal(ctx *gin.Context) {
	editForm := forms.JournalEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	journal := editForm.BindToModel()
	if err := journal.Update(); err != nil {
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

// MultiDelJournals
// @Summary 批量删除日志
// @Tags 日志
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/journals [delete]
func (j *JournalHandler) MultiDelJournals(ctx *gin.Context) {
	ids := ctx.Query("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请选择要删除的日志",
			Data: nil,
		})
		return
	}

	err := models.Journal{}.MultiDelByIds(ids)
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
