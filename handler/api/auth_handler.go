package api

import (
	"aries/config/setting"
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// @Summary 注册
// @Tags 授权
// @version 1.0
// @Accept application/json
// @Param regForm body form.RegisterForm true "注册表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/auth/register [post]
func Register(ctx *gin.Context) {
	regForm := form.RegisterForm{}
	result := util.Result{ // 定义 api 返回信息结构
		Code: util.Success,
		Msg:  "注册成功",
		Data: nil,
	}
	if err := ctx.ShouldBindJSON(&regForm); err != nil { // 表单校验失败
		result.Code = util.RequestError     // 请求数据有误
		result.Msg = util.GetFormError(err) // 获取表单错误信息
		ctx.JSON(http.StatusOK, result)     // 返回 json
		return
	}
	user := regForm.BindToModel() // 绑定表单数据到用户
	u, _ := user.GetByUsername()  // 根据用户名获取用户
	if u.Username != "" {         // 账号已被注册
		result.Code = util.RequestError
		result.Msg = "该用户已被注册"
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}
	if err := user.Create(); err != nil { // 创建用户 + 异常处理
		log.Errorln("error: ", err.Error())
		result.Code = util.ServerError
		result.Msg = "服务器端错误"
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}
	sysSetting := model.SysSetting{Name: "网站设置"}
	err := sysSetting.CreateOrUpdate()
	if err != nil {
		log.Errorln("error: ", err.Error())
		result.Code = util.ServerError
		result.Msg = "服务器端错误"
		ctx.JSON(http.StatusOK, result)
		return
	}
	item := model.SysSettingItem{
		SysId: sysSetting.ID,
		Key:   "site_url",
		Val:   regForm.SiteUrl,
	}
	var itemList []model.SysSettingItem
	itemList = append(itemList, item)
	err = model.SysSettingItem{}.MultiCreateOrUpdate(sysSetting.ID, itemList)
	if err != nil {
		log.Errorln("error: ", err.Error())
		result.Code = util.ServerError
		result.Msg = "服务器端错误"
		ctx.JSON(http.StatusOK, result)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// @Summary 登录
// @Tags 授权
// @version 1.0
// @Accept application/json
// @Param loginForm body form.LoginForm true "登录表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/auth/login [post]
func Login(ctx *gin.Context) {
	loginForm := form.LoginForm{}
	result := util.Result{ // 定义 api 返回信息结构
		Code: util.Success,
		Msg:  "登录成功",
		Data: nil,
	}
	if err := ctx.ShouldBindJSON(&loginForm); err != nil { // 表单校验失败
		result.Code = util.RequestError     // 请求数据有误
		result.Msg = util.GetFormError(err) // 获取表单错误信息
		ctx.JSON(http.StatusOK, result)     // 返回 json
		return
	}
	captchaConfig := &util.CaptchaConfig{
		Id:          loginForm.CaptchaId,
		VerifyValue: loginForm.CaptchaVal,
	}
	if !util.CaptchaVerify(captchaConfig) { // 校验验证码
		result.Code = util.RequestError // 请求数据有误
		result.Msg = "验证码错误"
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}
	user := loginForm.BindToModel() // 绑定表单数据到实体类
	u, _ := user.GetByUsername()    // 根据用户名获取用户
	if u.Username == "" {           // 用户不存在
		result.Code = util.RequestError
		result.Msg = "不存在该用户"
		ctx.JSON(http.StatusOK, result)
		return
	}
	if !util.VerifyPwd(u.Pwd, user.Pwd) { // 密码错误
		result.Code = util.RequestError
		result.Msg = "密码错误"
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}
	j := util.NewJWT()                             // 创建 JWT 实例
	token, err := j.CreateToken(util.CustomClaims{ // 生成 JWT token
		Username: u.Username,
		UserImg:  u.UserImg,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.
				Duration(setting.Config.Server.TokenExpireTime)).Unix(), // 设置过期时间
			IssuedAt: time.Now().Unix(),
		},
	})
	if err != nil { // 异常处理
		log.Errorln("error: ", err.Error())
		result.Code = util.ServerError
		result.Msg = "服务器端错误"
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = util.Token{ // 封装 Token 信息
		Token:    token,
		UserId:   u.ID,
		Username: u.Username,
		UserImg:  u.UserImg,
	}
	ctx.JSON(http.StatusOK, result) // 返回 json
}

// @Summary 创建验证码
// @Tags 授权
// @version 1.0
// @Accept application/json
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/auth/captcha [get]
func CreateCaptcha(ctx *gin.Context) {
	captcha := util.CaptchaConfig{} // 创建验证码配置结构
	result := util.Result{          // 返回数据结构
		Code: util.Success,
		Msg:  "验证码创建成功",
		Data: nil,
	}
	base64, err := util.GenerateCaptcha(&captcha) // 创建验证码
	if err != nil {                               // 异常处理
		result.Code = util.ServerError
		result.Msg = "服务器端错误"
		ctx.JSON(http.StatusOK, result)
		return
	}
	result.Data = gin.H{ // 封装 data
		"captcha_id":  captcha.Id,
		"captcha_url": base64,
	}
	ctx.JSON(http.StatusOK, result) // 返回 json 数据
}

// @Summary 忘记密码
// @Tags 授权
// @version 1.0
// @Accept application/json
// @Param forgetPwdForm body form.ForgetPwdForm true "忘记密码表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/auth/forget_pwd [post]
func ForgetPwd(ctx *gin.Context) {
	forgetPwdForm := form.ForgetPwdForm{}
	if err := ctx.ShouldBindJSON(&forgetPwdForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	user, _ := model.User{Email: forgetPwdForm.Email}.GetByEmail()
	if user.Username == "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "不存在该邮箱帐号",
			Data: nil,
		})
		return
	}
	verifyCode := ""
	_ = setting.Cache.Get(forgetPwdForm.Email, &verifyCode)
	if verifyCode == "" {
		verifyCode = util.CreateRandomCode(6)
		_ = setting.Cache.Set(forgetPwdForm.Email, verifyCode, time.Minute*15)
	}
	msg := gomail.NewMessage()
	// 设置收件人
	msg.SetHeader("To", forgetPwdForm.Email)
	// 设置发件人
	msg.SetAddressHeader("From", setting.Config.SMTP.Account, setting.Config.SMTP.Account)
	// 主题
	msg.SetHeader("Subject", "忘记密码验证")
	// 正文
	msg.SetBody("text/html", util.GetForgetPwdEmailHTML(user.Username, verifyCode))
	// 设置 SMTP 参数
	d := gomail.NewDialer(setting.Config.SMTP.Address, setting.Config.SMTP.Port,
		setting.Config.SMTP.Account, setting.Config.SMTP.Password)
	// 发送
	err := d.DialAndSend(msg)
	if err != nil {
		log.Error("验证码发送失败：", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "验证码发送失败，请检查 smtp 配置",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "验证码发送成功，请前往邮箱查看",
		Data: nil,
	})
}

// @Summary 重置密码
// @Tags 授权
// @version 1.0
// @Accept application/json
// @Param resetPwdForm body form.ResetPwdForm true "重置密码表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/auth/reset_pwd [post]
func ResetPwd(ctx *gin.Context) {
	resetPwdForm := form.ResetPwdForm{}
	if err := ctx.ShouldBindJSON(&resetPwdForm); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	verifyCode := ""
	_ = setting.Cache.Get(resetPwdForm.Email, &verifyCode)
	if verifyCode != resetPwdForm.VerifyCode {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "验证码无效或错误",
			Data: nil,
		})
		return
	}
	user := resetPwdForm.BindToModel()
	err := user.UpdatePwd()
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
		Msg:  "重置密码成功",
		Data: nil,
	})
}
