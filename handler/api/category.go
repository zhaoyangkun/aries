package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
