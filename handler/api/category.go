package api

import (
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取所有分类
func GetAllCategoriesHandler(ctx *gin.Context) {
	category := model.Category{}      // 建立 model 对象
	categoryList := category.GetAll() // 调用 model 对应方法，从数据库中获取所有分类
	result := util.Result{            // 封装返回体内容
		Code: http.StatusOK, // 状态码
		Msg:  "查询成功",        // 提示信息
		Data: categoryList,  // 数据
	}
	ctx.JSON(http.StatusOK, gin.H{ // 返回 json 数据
		"result": result,
	})
}
