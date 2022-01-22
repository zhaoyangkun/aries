package api

import (
	"aries/forms"
	"aries/handlers"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
}

// GetAllTags
// @Summary 获取所有标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_tags [get]
func (t *TagHandler) GetAllTags(ctx *gin.Context) {
	list, err := models.Tag{}.GetAllWithNoArticle()
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
		Msg:  "查询成功",
		Data: list,
	})
}

// GetTagsByPage
// @Summary 分页获取标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/tags [get]
func (t *TagHandler) GetTagsByPage(ctx *gin.Context) {
	pageForm := forms.TagPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)

	list, totalNum, err := models.Tag{}.GetByPage(&pageForm.Pagination, pageForm.Key)
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
		Msg:  "查询成功",
		Data: utils.GetPageData(list, totalNum, pageForm.Pagination),
	})
}

// GetTagById
// @Summary 根据 ID 获取标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param id path int true "ID"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/tags/{id} [get]
func (t *TagHandler) GetTagById(ctx *gin.Context) {
	id := ctx.Param("id")

	tag, err := models.Tag{}.GetById(id)
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
		Msg:  "查询成功",
		Data: tag,
	})
}

// AddTag
// @Summary 添加标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param addForm body forms.TagAddForm true "添加标签表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/tags [post]
func (t *TagHandler) AddTag(ctx *gin.Context) {
	addForm := forms.TagAddForm{}
	err := ctx.ShouldBindJSON(&addForm)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	// 校验标签名称唯一性
	existTag, _ := models.Tag{}.GetByName(addForm.Name)
	if existTag.Name != "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "已存在该标签，请勿重复添加",
			Data: nil,
		})
		return
	}

	tag := addForm.BindToModel()
	if err := tag.Create(); err != nil {
		log.Logger.Sugar().Error("Error: ", err.Error())
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
		Data: tag,
	})
}

// UpdateTag
// @Summary 修改标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param editForm body forms.TagEditForm true "修改标签表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/tags [put]
func (t *TagHandler) UpdateTag(ctx *gin.Context) {
	editForm := forms.TagEditForm{}
	err := ctx.ShouldBindJSON(&editForm)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	// 校验标签名称唯一性
	existTag, _ := models.Tag{}.GetByName(editForm.Name)
	if existTag.ID > 0 && existTag.ID != editForm.ID {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "标签名称不能重复",
			Data: nil,
		})
		return
	}

	tag := editForm.BindToModel()
	if err := tag.Update(); err != nil {
		log.Logger.Sugar().Error("Error: ", err.Error())
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
		Data: tag,
	})
}

// DeleteTag
// @Summary 删除标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param id path int true "ID"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/tags/{id} [delete]
func (t *TagHandler) DeleteTag(ctx *gin.Context) {
	id := ctx.Param("id")

	err := models.Tag{}.DeleteById(id)
	if err != nil {
		log.Logger.Sugar().Error("Error: ", err.Error())
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

// MultiDelTags
// @Summary 批量删除标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/tags [delete]
func (t *TagHandler) MultiDelTags(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "")
	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请求参数错误",
			Data: nil,
		})
		return
	}

	err := models.Tag{}.MultiDelByIds(ids)
	if err != nil {
		log.Logger.Sugar().Error("Error: ", err.Error())
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
