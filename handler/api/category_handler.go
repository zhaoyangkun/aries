package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary 获取所有分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param category_type query uint true "分类类型"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/all_categories [get]
func GetAllCategories(ctx *gin.Context) {
	categoryType := ctx.Query("category_type")
	cType, err := strconv.Atoi(categoryType)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请求参数有误",
			Data: nil,
		})
		return
	}
	category := model.Category{}                            // 建立 model 对象
	categoryList, err := category.GetAllByType(uint(cType)) // 调用 model 对应方法，从数据库中获取所有分类
	if err != nil {                                         // 异常处理
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		}) // 返回 json
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "查询成功",
		Data: categoryList,
	})
}

// @Summary 获取所有父级分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param category_type query uint true "分类类型"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/parent_categories [get]
func GetAllParentCategories(ctx *gin.Context) {
	categoryType := ctx.Query("category_type")
	cType, err := strconv.Atoi(categoryType)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请求参数有误",
			Data: nil,
		})
		return
	}
	category := model.Category{}
	categoryList, err := category.GetAllParents(uint(cType))
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
		Data: categoryList,
	})
}

// @Summary 分页获取分类数据
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Param category_type query uint true "分类类型"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/categories [get]
func GetCategoriesByPage(ctx *gin.Context) {
	pageForm := form.CategoryPageForm{}
	err := ctx.ShouldBindQuery(&pageForm)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	category := model.Category{}
	categoryList, totalNum, err := category.GetByPage(&pageForm.Pagination, pageForm.Key,
		*pageForm.CategoryType)
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
		Data: util.GetPageData(categoryList, totalNum, pageForm.Pagination),
	})
}

// @Summary 添加分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param addForm body form.CategoryAddForm true "添加分类表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/categories [post]
func AddCategory(ctx *gin.Context) {
	addForm := form.CategoryAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	// 校验分类名称唯一性
	existCategory, _ := model.Category{}.GetByName(addForm.Name)
	if existCategory.Name != "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "分类名不能重复",
			Data: nil,
		})
		return
	}
	// 校验分类 Url 唯一性
	existCategory, _ = model.Category{}.GetByUrl(addForm.Url)
	if existCategory.Url != "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "Url 不能重复",
			Data: nil,
		})
		return
	}
	category := addForm.BindToModel()
	if err := category.Create(); err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "创建成功",
		Data: category,
	})
}

// @Summary 修改分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param editForm body form.CategoryEditForm true "修改分类表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/categories [put]
func UpdateCategory(ctx *gin.Context) {
	editForm := form.CategoryEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	// 校验分类名称唯一性
	existCategory, _ := model.Category{}.GetByName(editForm.Name)
	if existCategory.ID > 0 && existCategory.ID != editForm.ID {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "分类名不能重复",
			Data: nil,
		})
		return
	}
	// 校验分类 Url 唯一性
	existCategory, _ = model.Category{}.GetByUrl(editForm.Url)
	if existCategory.ID > 0 && existCategory.ID != editForm.ID {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "Url 不能重复",
			Data: nil,
		})
		return
	}
	category := editForm.BindToModel()
	if err := category.Update(); err != nil {
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "修改成功",
		Data: category,
	})
}

// @Summary 删除分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/categories/{id} [delete]
func DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id")) // 将 string 转换为 int
	if err != nil {                          // 类型转换失败
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请求参数有误",
			Data: nil,
		})
		return
	}
	category := model.Category{}
	if err := category.DeleteById(uint(id)); err != nil { //删除分类，捕捉异常
		log.Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, util.Result{ // 删除成功
		Code: util.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// @Summary 批量删除分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/categories [delete]
func MultiDelCategories(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "") // 获取 ids
	if ids == "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}
	category := model.Category{}
	if err := category.MultiDelByIds(ids); err != nil {
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
