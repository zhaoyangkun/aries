package api

import (
	"aries/forms"
	"aries/handlers"
	"aries/log"
	"aries/models"
	"aries/utils"
	"github.com/gin-contrib/sessions"
	"io"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
}

// GetAllArticles
// @Summary 获取所有文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_articles [get]
func (a *ArticleHandler) GetAllArticles(ctx *gin.Context) {
	list, err := models.Article{}.GetAll()
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

// GetArticlesByPage
// @Summary 分页获取文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Param state query int false "状态"
// @Param category_id query int false "分类 ID"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles [get]
func (a *ArticleHandler) GetArticlesByPage(ctx *gin.Context) {
	pageForm := forms.ArticlePageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)

	list, totalNum, err := models.Article{}.GetByPage(&pageForm.Pagination, pageForm.Key,
		pageForm.State, pageForm.CategoryId)
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
		Data: utils.GetPageData(list, totalNum, pageForm.Pagination),
	})
}

// GetArticleById
// @Summary 根据 ID 获取文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles/{id} [get]
func (a *ArticleHandler) GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")

	article, err := models.Article{}.GetById(id)
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
		Data: article,
	})
}

// AddArticle
// @Summary 添加文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param addForm body forms.ArticleAddForm true "添加文章表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles [post]
func (a *ArticleHandler) AddArticle(ctx *gin.Context) {
	addForm := forms.ArticleAddForm{}
	if err := ctx.ShouldBindJSON(&addForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	existArticle, _ := models.Article{}.GetByUrl(addForm.URL)
	if existArticle.ID > 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "Url 不能重复",
			Data: nil,
		})
		return
	}

	article := addForm.BindToModel()
	if err := article.Create(addForm.TagIds); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "发布文章成功",
		Data: nil,
	})
}

// UpdateArticle
// @Summary 修改文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param editForm body forms.ArticleEditForm true "修改文章表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles [put]
func (a *ArticleHandler) UpdateArticle(ctx *gin.Context) {
	editForm := forms.ArticleEditForm{}
	if err := ctx.ShouldBindJSON(&editForm); err != nil {
		log.Logger.Sugar().Error("err: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	existArticle, _ := models.Article{}.GetByUrl(editForm.URL)
	if existArticle.ID > 0 && existArticle.ID != editForm.ID {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "Url 不能重复",
			Data: nil,
		})
		return
	}

	article := editForm.BindToModel()
	err := article.Update(editForm.TagIds)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "修改文章成功",
		Data: nil,
	})
}

// DeleteArticle
// @Summary 删除文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles/{id} [delete]
func (a *ArticleHandler) DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := models.Article{}.DeleteById(id)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// MultiDelArticles
// @Summary 批量删除文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles [delete]
func (a *ArticleHandler) MultiDelArticles(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "") // 获取 ids

	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}

	article := models.Article{}
	if err := article.MultiDelByIds(ids); err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	handlers.InitTmplVars()

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// ImportArticlesFromFiles
// @Summary 从文件导入文章
// @Tags 文章
// @version 1.0
// @Accept multipart/form-data
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles/files [post]
func (a *ArticleHandler) ImportArticlesFromFiles(ctx *gin.Context) {
	multiForm, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	files := multiForm.File["file[]"]

	if len(files) > 10 || len(files) == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "上传文件个数不能少于 1 个，也不能多于 10 个",
			Data: nil,
		})
		return
	}

	for _, file := range files {
		// 校验文件类型
		if path.Ext(file.Filename) != ".md" {
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.RequestError,
				Msg:  "只支持导入 md 格式的文件",
				Data: nil,
			})
			return
		}

		// 校验文件大小
		if file.Size > 2*1024*1024 {
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.RequestError,
				Msg:  "单个文件大小不能超过 2 MB",
				Data: nil,
			})
			return
		}
	}

	for _, file := range files {
		// 打开文件
		src, err := file.Open()
		if err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "读取文件内容失败",
				Data: nil,
			})
			return
		}

		// 关闭文件
		_ = src.Close()

		// 读取文件
		bytes, err := io.ReadAll(src)
		if err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}

		article := models.Article{
			Content: string(bytes),
			Title:   utils.GetFileNameOnly(file.Filename),
		}
		err = article.SaveFromFile()
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
		Msg:  "导入成功",
		Data: nil,
	})
}

// RecycleOrRecoverArticle
// @Summary 回收或恢复文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/articles/recycle/{id} [patch]
func (a *ArticleHandler) RecycleOrRecoverArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := models.Article{}.RecycleOrRecover(id)
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
		Msg:  "操作成功",
		Data: nil,
	})
}

// MoveArticleUp
// @Summary 向上移动文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param oderForm body forms.ArticleOrderForm true "文章排序表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/articles/up [patch]
func (a *ArticleHandler) MoveArticleUp(ctx *gin.Context) {
	orderForm := forms.ArticleOrderForm{}
	if err := ctx.ShouldBindJSON(&orderForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	currArticle := orderForm.BindToModel()
	preArticle, _ := currArticle.GetPrevious(currArticle.OrderId, currArticle.IsTop, false)
	if preArticle.ID == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "无法再向上移动",
			Data: nil,
		})
		return
	}

	err := models.Article{}.MoveUp(currArticle.ID, preArticle.ID, currArticle.OrderId, preArticle.OrderId)
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
		Msg:  "向上移动成功",
		Data: nil,
	})
}

// MoveArticleDown
// @Summary 向下移动文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param oderForm body forms.ArticleOrderForm true "文章排序表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles/down [patch]
func (a *ArticleHandler) MoveArticleDown(ctx *gin.Context) {
	orderForm := forms.ArticleOrderForm{}
	if err := ctx.ShouldBindJSON(&orderForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	currArticle := orderForm.BindToModel()
	nextArticle, _ := currArticle.GetNext(currArticle.OrderId, currArticle.IsTop, false)
	if nextArticle.ID == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "无法再向下移动",
			Data: nil,
		})
		return
	}

	err := models.Article{}.MoveDown(currArticle.ID, nextArticle.ID, currArticle.OrderId, nextArticle.OrderId)
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
		Msg:  "向下移动成功",
		Data: nil,
	})
}

// CheckArticlePwd
// @Summary 校验文章密码
// @Tags 文章
// @version 1.0
// @Accept application/x-www-form-urlencoded
// @Param oderForm body forms.ArticlePwdForm true "文章密码表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/articles/check [post]
func (a *ArticleHandler) CheckArticlePwd(ctx *gin.Context) {
	pwdForm := forms.ArticlePwdForm{}
	if err := ctx.ShouldBind(&pwdForm); err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  utils.GetFormError(err),
			Data: nil,
		})
		return
	}

	article, err := models.Article{}.GetById(pwdForm.ArticleId)
	if err != nil {
		log.Logger.Sugar().Error("error: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	if utils.VerifyPwd(article.Pwd, pwdForm.Pwd) {
		session := sessions.Default(ctx)
		session.Set("article-"+pwdForm.ArticleId, true)
		_ = session.Save()
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.Success,
			Msg:  "OK",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.RequestError,
		Msg:  "密码错误",
		Data: nil,
	})
}
