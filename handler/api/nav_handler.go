package api

import (
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NavHandler struct {
}

// 获取所有菜单
func (n *NavHandler) GetAllNavs(ctx *gin.Context) {
	list, err := model.Nav{}.GetAll()
	if err != nil {
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
