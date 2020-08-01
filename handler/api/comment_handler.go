package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// @Summary 获取所有评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/all_comments [get]
func GetAllComments(ctx *gin.Context) {
	list, err := model.Comment{}.GetAll()
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
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
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/comments [get]
func GetCommentsByPage(ctx *gin.Context) {
	pageForm := form.CommentPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)
	list, total, err := model.Comment{}.GetByPage(&pageForm.Pagination, pageForm.Key, pageForm.Type, pageForm.State)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "查询成功",
		Data: util.GetPageData(list, total, pageForm.Pagination),
	})
}

// @Summary 发表评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param form body form.CommentAddForm false "添加评论表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/comments [post]
func AddComment(ctx *gin.Context) {
	addForm := form.CommentAddForm{}
	err := ctx.ShouldBindJSON(&addForm)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	comment := addForm.BindToModel()
	if err := comment.Create(); err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "发表评论成功",
		Data: nil,
	})
}

// @Summary 修改评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param form body form.CommentEditForm false "修改评论表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/comments [put]
func UpdateComment(ctx *gin.Context) {
	editForm := form.CommentEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	comment := editForm.BindToModel()
	if err := comment.Update(); err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "修改评论成功",
		Data: nil,
	})
}

// @Summary 删除评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param id path uint true "id"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/comments/{id} [delete]
func DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	err := model.Comment{}.DeleteById(id)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// @Summary 批量删除评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/comments [delete]
func MultiDelComments(ctx *gin.Context) {
	ids := ctx.Query("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}
	err := model.Comment{}.MultiDelByIds(ids)
	if err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}
