package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary 获取所有标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/all_tags [get]
func GetAllTags(ctx *gin.Context) {
	list, err := model.Tag{}.GetAll()
	if err != nil {
		log.Println("Error: ", err.Error())
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

// @Summary 分页获取标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/tags [get]
func GetTagsByPage(ctx *gin.Context) {
	pageForm := form.TagPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)
	list, totalNum, err := model.Tag{}.GetByPage(&pageForm.Pagination, pageForm.Key)
	if err != nil {
		log.Println("Error: ", err.Error())
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
		Data: util.GetPageData(list, totalNum, pageForm.Pagination),
	})
}

// @Summary 根据 ID 获取标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param id path int true "ID"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/tags/{id} [get]
func GetTagById(ctx *gin.Context) {
	id := ctx.Param("id")
	tag, err := model.Tag{}.GetById(id)
	if err != nil {
		log.Println("Error: ", err.Error())
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
		Data: tag,
	})
}

// @Summary 添加标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param addForm body form.TagAddForm true "添加标签表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/tags [post]
func AddTag(ctx *gin.Context) {
	addForm := form.TagAddForm{}
	err := ctx.ShouldBindJSON(&addForm)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	// 校验标签名称唯一性
	existTag, _ := model.Tag{}.GetByName(addForm.Name)
	if existTag.Name != "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "已存在该标签，请勿重复添加",
			Data: nil,
		})
		return
	}
	tag := addForm.BindToModel()
	if err := tag.Create(); err != nil {
		log.Println("Error: ", err.Error())
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
		Data: tag,
	})
}

// @Summary 修改标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param editForm body form.TagEditForm true "修改标签表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/tags [put]
func UpdateTag(ctx *gin.Context) {
	editForm := form.TagEditForm{}
	err := ctx.ShouldBindJSON(&editForm)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	// 校验标签名称唯一性
	existTag, _ := model.Tag{}.GetByName(editForm.Name)
	if existTag.ID > 0 && existTag.ID != editForm.ID {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "标签名称不能重复",
			Data: nil,
		})
		return
	}
	tag := editForm.BindToModel()
	if err := tag.Update(); err != nil {
		log.Println("Error: ", err.Error())
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
		Data: tag,
	})
}

// @Summary 删除标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param id path int true "ID"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(ctx *gin.Context) {
	id := ctx.Param("id")
	err := model.Tag{}.DeleteById(id)
	if err != nil {
		log.Println("Error: ", err.Error())
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

// @Summary 批量删除标签
// @Tags 标签
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/tags [delete]
func MultiDelTags(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "")
	if ids == "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请求参数错误",
			Data: nil,
		})
		return
	}
	err := model.Tag{}.MultiDelByIds(ids)
	if err != nil {
		log.Println("Error: ", err.Error())
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
