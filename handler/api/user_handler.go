package api

import (
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary 获取所有用户
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/all_users [get]
func GetAllUsers(ctx *gin.Context) {
	list, err := model.User{}.GetAll()
	if err != nil {
		log.Println("数据库错误：", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务端错误",
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
