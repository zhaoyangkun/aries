package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
		Data: list,
	})
}

// @Summary 更新用户信息
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Param userForm body form.UserInfoForm true "用户信息表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/users [put]
func UpdateUser(ctx *gin.Context) {
	userForm := form.UserInfoForm{}
	if err := ctx.ShouldBindJSON(&userForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	user := userForm.BindToModel()
	if err := user.Update(); err != nil {
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
		Msg:  "修改成功",
		Data: nil,
	})
}
