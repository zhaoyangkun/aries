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

type CategoryHandler struct {
}

// GetAllCategories
// @Summary 获取所有分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param category_type query uint true "分类类型"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_categories [get]
func (c *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	categoryType := ctx.Query("category_type")

	cType, err := strconv.Atoi(categoryType)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请求参数有误",
			Data: nil,
		})
		return
	}

	category := models.Category{}                           // 建立 model 对象
	categoryList, err := category.GetAllByType(uint(cType)) // 调用 model 对应方法，从数据库中获取所有分类
	if err != nil {                                         // 异常处理
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		}) // 返回 json
		return
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "查询成功",
		Data: categoryList,
	})
}

// GetAllParentCategories
// @Summary 获取所有父级分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param category_type query uint true "分类类型"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/parent_categories [get]
func (c *CategoryHandler) GetAllParentCategories(ctx *gin.Context) {
	categoryType := ctx.Query("category_type")

	cType, err := strconv.Atoi(categoryType)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请求参数有误",
			Data: nil,
		})
		return
	}

	category := models.Category{}
	categoryList, err := category.GetAllParents(uint(cType))
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
		Data: categoryList,
	})
}

// GetCategoriesByPage
// @Summary 分页获取分类数据
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Param category_type query uint true "分类类型"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories [get]
func (c *CategoryHandler) GetCategoriesByPage(ctx *gin.Context) {
	pageForm := forms.CategoryPageForm{}
	if err := ctx.ShouldBindQuery(&pageForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	category := models.Category{}
	categoryList, totalNum, err := category.GetByPage(&pageForm.Pagination, pageForm.Key,
		*pageForm.CategoryType)
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
		Data: utils.GetPageData(categoryList, totalNum, pageForm.Pagination),
	})
}

// AddArticleCategory
// @Summary 添加文章分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param addForm body forms.ArticleCategoryAddForm true "添加文章分类表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories/article [post]
func (c *CategoryHandler) AddArticleCategory(ctx *gin.Context) {
	addForm := forms.ArticleCategoryAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	// 校验分类名称唯一性
	existCategory, _ := models.Category{}.GetByName(addForm.Name, addForm.Type)
	if existCategory.Name != "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "分类名不能重复",
			Data: nil,
		})
		return
	}

	// 校验分类 Url 唯一性
	existCategory, _ = models.Category{}.GetByUrl(addForm.Url)
	if existCategory.Url != "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "Url 不能重复",
			Data: nil,
		})
		return
	}

	category := addForm.BindToModel()
	if err := category.Create(); err != nil {
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
		Msg:  "创建成功",
		Data: category,
	})
}

// UpdateArticleCategory
// @Summary 修改文章分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param editForm body forms.ArticleCategoryEditForm true "修改文章分类表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories/article [put]
func (c *CategoryHandler) UpdateArticleCategory(ctx *gin.Context) {
	editForm := forms.ArticleCategoryEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	// 校验分类名称唯一性
	existCategory, _ := models.Category{}.GetByName(editForm.Name, editForm.Type)
	if existCategory.ID > 0 && existCategory.ID != editForm.ID {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "分类名不能重复",
			Data: nil,
		})
		return
	}

	// 校验分类 Url 唯一性
	existCategory, _ = models.Category{}.GetByUrl(editForm.Url)
	if existCategory.ID > 0 && existCategory.ID != editForm.ID {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "Url 不能重复",
			Data: nil,
		})
		return
	}

	category := editForm.BindToModel()
	if err := category.Update(); err != nil {
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
		Data: category,
	})
}

// AddLinkCategory
// @Summary 添加友链分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param addForm body forms.LinkCategoryAddForm true "添加友链分类表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories/link [post]
func (c *CategoryHandler) AddLinkCategory(ctx *gin.Context) {
	addForm := forms.LinkCategoryAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	// 校验分类名称唯一性
	existCategory, _ := models.Category{}.GetByName(addForm.Name, addForm.Type)
	if existCategory.Name != "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "分类名不能重复",
			Data: nil,
		})
		return
	}

	category := addForm.BindToModel()
	if err := category.Create(); err != nil {
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
		Msg:  "创建成功",
		Data: category,
	})
}

// UpdateLinkCategory
// @Summary 修改友链分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param editForm body forms.LinkCategoryEditForm true "修改友链分类表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories/link [put]
func (c *CategoryHandler) UpdateLinkCategory(ctx *gin.Context) {
	editForm := forms.LinkCategoryEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	// 校验分类名称唯一性
	existCategory, _ := models.Category{}.GetByName(editForm.Name, editForm.Type)
	if existCategory.ID > 0 && existCategory.ID != editForm.ID {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "分类名不能重复",
			Data: nil,
		})
		return
	}

	category := editForm.BindToModel()
	if err := category.Update(); err != nil {
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
		Data: category,
	})
}

// AddGalleryCategory
// @Summary 添加图库分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param addForm body forms.GalleryCategoryAddForm true "添加图库分类表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories/gallery [post]
func (c *CategoryHandler) AddGalleryCategory(ctx *gin.Context) {
	addForm := forms.GalleryCategoryAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	// 校验分类名称唯一性
	existCategory, _ := models.Category{}.GetByName(addForm.Name, addForm.Type)
	if existCategory.Name != "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "分类名不能重复",
			Data: nil,
		})
		return
	}

	category := addForm.BindToModel()
	if err := category.Create(); err != nil {
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
		Msg:  "创建成功",
		Data: category,
	})
}

// UpdateGalleryCategory
// @Summary 修改图库分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param editForm body forms.GalleryCategoryEditForm true "修改图库分类表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories/gallery [put]
func (c *CategoryHandler) UpdateGalleryCategory(ctx *gin.Context) {
	editForm := forms.GalleryCategoryEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	// 校验分类名称唯一性
	existCategory, _ := models.Category{}.GetByName(editForm.Name, editForm.Type)
	if existCategory.ID > 0 && existCategory.ID != editForm.ID {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "分类名不能重复",
			Data: nil,
		})
		return
	}

	category := editForm.BindToModel()
	if err := category.Update(); err != nil {
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
		Data: category,
	})
}

// DeleteCategory
// @Summary 删除分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories/{id} [delete]
func (c *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id")) // 将 string 转换为 int

	if err != nil { // 类型转换失败
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请求参数有误",
			Data: nil,
		})
		return
	}

	category := models.Category{}
	if err := category.DeleteById(uint(id)); err != nil { //删除分类，捕捉异常
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{ // 删除成功
		Code: utils.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// MultiDelCategories
// @Summary 批量删除分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/categories [delete]
func (c *CategoryHandler) MultiDelCategories(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "") // 获取 ids

	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}

	category := models.Category{}
	if err := category.MultiDelByIds(ids); err != nil {
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
