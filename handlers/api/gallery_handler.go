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

type GalleryHandler struct {
}

// GetAllGalleries
// @Summary 获取所有图库
// @Tags 图库
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_galleries [get]
func (g *GalleryHandler) GetAllGalleries(ctx *gin.Context) {
	list, err := models.Gallery{}.GetAll()
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

// GetGalleryById
// @Summary 根据 ID 获取图库
// @Tags 图库
// @version 1.0
// @Accept application/json
// @Param id path uint true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/galleries/{id} [get]
func (g *GalleryHandler) GetGalleryById(ctx *gin.Context) {
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

	obj, err := models.Gallery{}.GetById(uint(intId))
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

// GetGalleriesByPage
// @Summary 分页获取图库
// @Tags 图库
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param category_id query uint false "分类 ID"
// @Param key query string false "关键词"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/galleries [get]
func (g *GalleryHandler) GetGalleriesByPage(ctx *gin.Context) {
	pageForm := forms.GalleryPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)

	list, total, err := models.Gallery{}.GetByPage(&pageForm.Pagination, pageForm.CategoryId, pageForm.Key)
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

// CreateGallery
// @Summary 创建图库
// @Tags 图库
// @version 1.0
// @Accept application/json
// @Param addForm body forms.AddGalleryForm false "创建图库表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/galleries [post]
func (g *GalleryHandler) CreateGallery(ctx *gin.Context) {
	addForm := forms.AddGalleryForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	gallery := addForm.BindToModel()
	if err := gallery.Create(); err != nil {
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

// UpdateGallery
// @Summary 修改图库
// @Tags 图库
// @version 1.0
// @Accept application/json
// @Param editForm body forms.EditGalleryForm false "修改图库表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/galleries [put]
func (g *GalleryHandler) UpdateGallery(ctx *gin.Context) {
	editForm := forms.EditGalleryForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	gallery := editForm.BindToModel()
	if err := gallery.Update(); err != nil {
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

// MultiDelGalleries
// @Summary 批量删除图库
// @Tags 图库
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/galleries [delete]
func (g *GalleryHandler) MultiDelGalleries(ctx *gin.Context) {
	ids := ctx.Query("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请选择要删除的日志",
			Data: nil,
		})
		return
	}

	err := models.Gallery{}.MultiDelByIds(ids)
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
