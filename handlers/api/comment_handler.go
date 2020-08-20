package api

import (
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
}

// @Summary 获取所有评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_comments [get]
func (c *CommentHandler) GetAllComments(ctx *gin.Context) {
	list, err := models.Comment{}.GetAll()
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

// @Summary 分页获取评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Param type query string false "类型"
// @Param state query uint false "状态"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments [get]
func (c *CommentHandler) GetCommentsByPage(ctx *gin.Context) {
	pageForm := forms.CommentPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)

	list, total, err := models.Comment{}.GetByPage(&pageForm.Pagination, pageForm.Key, pageForm.Type, pageForm.State)
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

// @Summary 发表评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param form body forms.CommentAddForm false "添加评论表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments [post]
func (c *CommentHandler) AddComment(ctx *gin.Context) {
	addForm := forms.CommentAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	comment := addForm.BindToModel()
	if err := comment.Create(); err != nil {
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
		Msg:  "发表评论成功",
		Data: nil,
	})
}

// @Summary 修改评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param form body forms.CommentEditForm false "修改评论表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments [put]
func (c *CommentHandler) UpdateComment(ctx *gin.Context) {
	editForm := forms.CommentEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	comment := editForm.BindToModel()
	if err := comment.Update(); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "修改评论成功",
		Data: nil,
	})
}

// @Summary 删除评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param id path uint true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments/{id} [delete]
func (c *CommentHandler) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")

	err := models.Comment{}.DeleteById(id)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// @Summary 批量删除评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments [delete]
func (c *CommentHandler) MultiDelComments(ctx *gin.Context) {
	ids := ctx.Query("ids")

	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}

	err := models.Comment{}.MultiDelByIds(ids)
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
