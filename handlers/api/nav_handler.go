package api

import (
	"aries/models"
	"aries/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NavHandler struct {
}

// 获取所有菜单
func (n *NavHandler) GetAllNavs(ctx *gin.Context) {
	list, err := models.Nav{}.GetAll()
	if err != nil {
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
