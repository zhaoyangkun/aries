package api

import (
	"aries/config"
	"aries/form"
	"aries/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 注册
func Register(ctx *gin.Context) {
	regForm := form.RegisterForm{}
	result := util.Result{ // 定义 api 返回信息结构
		Code: util.Success,
		Msg:  "注册成功",
		Data: nil,
	}
	if err := ctx.ShouldBindJSON(&regForm); err == nil { // 表单校验通过
		user := regForm.BindToModel() // 绑定表单数据到用户
		u, _ := user.GetByUsername()  // 根据用户名获取用户
		if u.Username != "" {         // 账号已被注册
			result.Code = util.RequestError
			result.Msg = "该用户已被注册"
		} else {
			if err := user.Create(); err != nil { // 创建用户 + 异常处理
				log.Println(err)
				result.Code = util.ServerError
				result.Msg = "服务器内部错误"
			}
		}
	} else { // 表单校验失败
		log.Println("form: ", regForm)
		result.Code = util.RequestError     // 请求数据有误
		result.Msg = util.GetFormError(err) // 获取表单错误信息
	}
	ctx.JSON(http.StatusOK, result) // 返回 json
}

// 登录
func Login(ctx *gin.Context) {
	loginForm := form.LoginForm{}
	result := util.Result{ // 定义 api 返回信息结构
		Code: util.Success,
		Msg:  "登录成功",
		Data: nil,
	}
	if err := ctx.ShouldBindJSON(&loginForm); err == nil { // 表单校验通过
		user := loginForm.BindToModel()                           // 绑定表单数据到实体类
		u, _ := user.GetByUsername()                              // 根据用户名获取用户
		if u.Username == "" || !util.VerifyPwd(u.Pwd, user.Pwd) { // 用户名或密码错误
			result.Code = util.RequestError
			result.Msg = "用户名或密码错误"
		} else { // 登录成功
			j := util.NewJWT()                             // 创建 JWT 实例
			token, err := j.CreateToken(util.CustomClaims{ // 生成 JWT token
				Username: u.Username,
				UserImg:  u.UserImg,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Second * time.
						Duration(config.AppConfig.TokenExpireTime)).Unix(), // 设置过期时间
					IssuedAt: time.Now().Unix(),
				},
			})
			if err != nil { // 异常处理
				log.Println(err.Error())
				result.Code = util.ServerError
				result.Msg = "服务器内部错误"
			} else { // 封装 Token 信息
				result.Data = util.Token{
					Token:    token,
					Username: u.Username,
					UserImg:  u.UserImg,
				}
			}
		}
	} else { // 表单校验失败
		result.Code = util.RequestError     // 请求数据有误
		result.Msg = util.GetFormError(err) // 获取表单错误信息
	}
	ctx.JSON(http.StatusOK, result) // 返回 json
}
