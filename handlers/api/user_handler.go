package api

import (
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

// GetAllUsers
// @Summary 获取所有用户
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_users [get]
func (u *UserHandler) GetAllUsers(ctx *gin.Context) {
	list, err := models.User{}.GetAll()
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

// UpdateUser
// @Summary 更新用户信息
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Param userForm body forms.UserInfoForm true "用户信息表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/users [put]
func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	userForm := forms.UserInfoForm{}
	if err := ctx.ShouldBindJSON(&userForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	user := userForm.BindToModel()
	if err := user.Update(); err != nil {
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
		Data: nil,
	})
}

// UpdateUserPwd
// @Summary 修改密码
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Param pwdForm body forms.PwdForm true "修改密码表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/users/pwd [put]
func (u *UserHandler) UpdateUserPwd(ctx *gin.Context) {
	pwdForm := forms.PwdForm{}
	if err := ctx.ShouldBindJSON(&pwdForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	oldUser, err := models.User{Username: pwdForm.Username}.GetByUsername()
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	if !utils.VerifyPwd(oldUser.Pwd, pwdForm.OldPwd) {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "旧密码错误",
			Data: nil,
		})
		return
	}

	oldUser.Pwd = pwdForm.NewPwd
	err = oldUser.UpdatePwd()
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
		Msg:  "修改密码成功，请重新登录",
		Data: nil,
	})
}
