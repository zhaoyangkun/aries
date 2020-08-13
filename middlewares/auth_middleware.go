package middlewares

import (
	"aries/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//  JWT 权限校验中间件
func JWTAuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token") // 从请求体头部获取 token
		result := utils.Result{                      // 封装返回体内容
			Code: utils.Forbidden, // 状态码
			Msg:  "",              // 提示信息
			Data: nil,             // 数据
		}
		if token == "" { // token 为空
			result.Msg = "请求未携带 Token，无权访问"
			context.JSON(http.StatusOK, result) // 返回 json
			context.Abort()                     // 停止处理 handler
			return
		}
		jwt := utils.NewJWT()                // 创建新的 JWT 实例
		claims, err := jwt.ParseToken(token) // 解析 Token
		if err != nil {                      // 错误处理
			result.Msg = err.Error()
			context.JSON(http.StatusOK, result) // 返回 json
			context.Abort()
			return
		}
		// 继续交由下一个 handler 处理,并将解析出的信息传递下去
		context.Set("claims", claims)
	}
}