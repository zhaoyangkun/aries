package api

import (
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 获取所有分类
func GetAllCategories(ctx *gin.Context) {
	category := model.Category{}           // 建立 model 对象
	categoryList, err := category.GetAll() // 调用 model 对应方法，从数据库中获取所有分类
	result := util.Result{                 // 封装返回体内容
		Code: http.StatusOK, // 状态码
		Msg:  "查询成功",        // 提示信息
		Data: categoryList,  // 数据
	}
	if err != nil { // 异常处理
		log.Println(err.Error())
		result.Code = http.StatusInternalServerError
		result.Msg = "服务器内部错误"
		result.Data = nil
	}
	ctx.JSON(result.Code, result) // 返回 json
}
