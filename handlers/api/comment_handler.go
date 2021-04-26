package api

import (
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

type CommentHandler struct {
}

// GetAllComments
// @Summary 获取所有评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_comments [get]
func (c *CommentHandler) GetAllComments(ctx *gin.Context) {
	list, err := models.Comment{}.GetAll()
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

// GetCommentsByPage
// @Summary 分页获取评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Param type query string false "类型"
// @Param state query uint false "状态"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments [get]
func (c *CommentHandler) GetCommentsByPage(ctx *gin.Context) {
	pageForm := forms.CommentPageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)

	list, total, err := models.Comment{}.GetByPage(&pageForm.Pagination, pageForm.Key,
		pageForm.ArticleId, pageForm.PageId, pageForm.Type, pageForm.State, pageForm.IsParent)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	disNum := uint(0)
	if pageForm.IsParent == 1 {
		disNum, err = models.Comment{}.GetDisCount(pageForm.PageId, pageForm.ArticleId)
		if err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "查询成功",
		Data: gin.H{
			"ok":          true,                           // 是否成功
			"data":        list,                           // 分页数据
			"total_num":   total,                          // 总条数
			"total_pages": pageForm.Pagination.TotalPages, // 总页数
			"page":        pageForm.Pagination.Page,       // 页码
			"size":        pageForm.Pagination.Size,       // 每页条数
			"dis_num":     disNum,
		},
	})
}

// AddComment
// @Summary 发表评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param form body forms.CommentAddForm false "添加评论表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments [post]
func (c *CommentHandler) AddComment(ctx *gin.Context) {
	addForm := forms.CommentAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	comment := addForm.BindToModel()
	if err := comment.Create(); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	commentItems, _ := models.SysSettingItem{}.GetBySysSettingName("评论设置")
	siteItems, _ := models.SysSettingItem{}.GetBySysSettingName("网站设置")
	emailSetting, _ := models.SysSettingItem{}.GetBySysSettingName("邮件设置")
	if len(emailSetting) == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "邮件发送失败，请先配置 SMTP",
			Data: nil,
		})
		return
	}
	userList, _ := models.User{}.GetAll()
	// 若开启邮件回复功能，发送回复邮件
	if isReplyOn, ok := commentItems["is_reply_on"]; ok && isReplyOn == "1" {
		msg := gomail.NewMessage()
		if comment.ParentCommentId == 0 {
			// 设置收件人
			msg.SetHeader("To", userList[0].Email)
			// 设置发件人
			msg.SetAddressHeader("From", emailSetting["account"], emailSetting["account"])
			// 主题
			msg.SetHeader("Subject", "评论通知")
			// 正文
			if comment.ArticleId > 0 {
				article, _ := models.Article{}.GetById(strconv.Itoa(int(comment.ArticleId)))
				msg.SetBody("text/html", utils.GetCommentEmailHTML(
					siteItems["site_name"], siteItems["site_url"], "评论通知",
					userList[0].Nickname, comment.NickName, comment.Url, article.Title,
					siteItems["site_url"]+"/articles/"+article.URL, comment.MDContent,
				))
			} else {
				page, _ := models.Page{}.GetById(comment.PageId)
				msg.SetBody("text/html", utils.GetCommentEmailHTML(
					siteItems["site_name"], siteItems["site_url"], "评论通知",
					userList[0].Nickname, comment.NickName, comment.Url, page.Title,
					siteItems["site_url"]+"/custom/"+page.Url, comment.MDContent,
				))
			}
		} else {
			parentComment, _ := models.Comment{}.GetById(comment.ParentCommentId)
			// 设置收件人
			msg.SetHeader("To", parentComment.Email)
			// 设置发件人
			msg.SetAddressHeader("From", emailSetting["account"], emailSetting["account"])
			// 主题
			msg.SetHeader("Subject", "评论通知")
			if comment.ArticleId > 0 {
				article, _ := models.Article{}.GetById(strconv.Itoa(int(comment.ArticleId)))
				msg.SetBody("text/html", utils.GetReplyEmailHTML(
					siteItems["site_name"], siteItems["site_url"], "评论通知",
					parentComment.NickName, comment.NickName, comment.Url, article.Title,
					siteItems["site_url"]+"/articles/"+article.URL, comment.MDContent,
				))
			} else {
				page, _ := models.Page{}.GetById(comment.PageId)
				msg.SetBody("text/html", utils.GetReplyEmailHTML(
					siteItems["site_name"], siteItems["site_url"], "评论通知",
					parentComment.NickName, comment.NickName, comment.Url, page.Title,
					siteItems["site_url"]+"/custom/"+page.Url, comment.MDContent,
				))
			}
		}
		// 设置 SMTP 参数
		port, _ := strconv.Atoi(emailSetting["port"])
		// 设置 SMTP 参数
		d := gomail.NewDialer(emailSetting["address"], port, emailSetting["account"], emailSetting["pwd"])
		// 发送邮件
		err := d.DialAndSend(msg)
		if err != nil {
			log.Logger.Sugar().Error("回复邮件发送失败：", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "回复邮件发送失败，请检查 smtp 配置",
				Data: nil,
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "发表评论成功",
		Data: nil,
	})
}

// UpdateComment
// @Summary 修改评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param form body forms.CommentEditForm false "修改评论表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments [put]
func (c *CommentHandler) UpdateComment(ctx *gin.Context) {
	editForm := forms.CommentEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	comment := editForm.BindToModel()
	if err := comment.Update(); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}
	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "修改评论成功",
		Data: nil,
	})
}

// DeleteComment
// @Summary 删除评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param id path uint true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments/{id} [delete]
func (c *CommentHandler) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")

	err := models.Comment{}.DeleteById(id)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// MultiDelComments
// @Summary 批量删除评论
// @Tags 评论
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/comments [delete]
func (c *CommentHandler) MultiDelComments(ctx *gin.Context) {
	ids := ctx.Query("ids")

	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}

	err := models.Comment{}.MultiDelByIds(ids)
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
		Msg:  "删除成功",
		Data: nil,
	})
}
