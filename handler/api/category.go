package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
)

// @Summary 获取所有分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/categories/all [get]
func GetAllCategories(ctx *gin.Context) {
	category := model.Category{}           // 建立 model 对象
	categoryList, err := category.GetAll() // 调用 model 对应方法，从数据库中获取所有分类
	result := util.Result{                 // 封装返回体内容
		Code: util.Success, // 状态码
		Msg:  "查询成功",       // 提示信息
		Data: categoryList, // 数据
	}
	if err != nil { // 异常处理
		log.Println(err.Error())
		result.Code = util.ServerError
		result.Msg = "服务器内部错误"
		result.Data = nil
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}
	ctx.JSON(http.StatusOK, result) // 返回 json
}

// @Summary 分页获取分类数据
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/categories [get]
func GetCategoriesByPage(ctx *gin.Context) {
	pageForm := form.CategoryPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)
	category := model.Category{}
	categoryList, totalNum, err := category.GetByPage(&pageForm.Pagination, pageForm.Key)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器错误",
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
	category := addForm.BindToModel()
	if err := category.Create(); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器内部错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "创建成功",
		Data: nil,
	})
}

// @Summary 修改分类
// @Tags 分类
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Param editForm body form.CategoryEditForm true "修改分类表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/categories/{id} [put]
func UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id")) // 将 string 转换为 int
	if err != nil {                          // 类型转换失败
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请求参数有误",
			Data: nil,
		})
		return
	}
	editForm := form.CategoryEditForm{Id: uint(id)}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	category := editForm.BindToModel()
	if err := category.Update(); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器内部错误",
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
	category := model.Category{
		Model: gorm.Model{ID: uint(id)},
	}
	if err := category.DeleteById(); err != nil { //删除分类，捕捉异常
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器内部错误",
			Data: nil,
		})
		return
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
func MultiDeleteCategory(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "") // 获取 ids
	category := model.Category{}
	if err := category.MultiDelByIds(ids); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器内部错误",
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
