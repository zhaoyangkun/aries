package api

import (
	"aries/form"
	"aries/model"
	"aries/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

// @Summary 获取所有文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/all_articles [get]
func GetAllArticles(ctx *gin.Context) {
	list, err := model.Article{}.GetAll()
	if err != nil {
		log.Println("error: ", err.Error())
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

// @Summary 分页获取文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Param state query int false "状态"
// @Param category_id query int false "分类 ID"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/articles [get]
func GetArticlesByPage(ctx *gin.Context) {
	pageForm := form.ArticlePageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)
	list, totalNum, err := model.Article{}.GetByPage(&pageForm.Pagination, pageForm.Key,
		pageForm.State, pageForm.CategoryId)
	if err != nil {
		log.Println("error: ", err.Error())
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
		Data: util.GetPageData(list, totalNum, pageForm.Pagination),
	})
}

// @Summary 根据 ID 获取文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/articles/{id} [get]
func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	article, err := model.Article{}.GetById(id)
	if err != nil {
		log.Println("error: ", err.Error())
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
		Data: article,
	})
}

// @Summary 添加文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param addForm body form.ArticleAddForm true "添加文章表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/articles [post]
func AddArticle(ctx *gin.Context) {
	addForm := form.ArticleAddForm{}
	err := ctx.ShouldBindJSON(&addForm)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	article := addForm.BindToModel()
	err = article.Create(addForm.TagIds)
	if err != nil {
		log.Println("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "发布文章成功",
		Data: nil,
	})
}

// @Summary 修改文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param editForm body form.ArticleEditForm true "修改文章表单"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/articles [put]
func UpdateArticle(ctx *gin.Context) {
	editForm := form.ArticleEditForm{}
	err := ctx.ShouldBindJSON(&editForm)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  util.GetFormError(err),
			Data: nil,
		})
		return
	}
	article := editForm.BindToModel()
	err = article.Update(editForm.TagIds)
	if err != nil {
		log.Println("error: ", err.Error())
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "修改文章成功",
		Data: nil,
	})
}

// @Summary 删除文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param id path int true "id"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	err := model.Article{}.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器内部错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// @Summary 批量删除文章
// @Tags 文章
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/articles [delete]
func MultiDelArticles(ctx *gin.Context) {
	ids := ctx.DefaultQuery("ids", "") // 获取 ids
	if ids == "" {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}
	article := model.Article{}
	if err := article.MultiDelByIds(ids); err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// @Summary 从文件导入文章
// @Tags 文章
// @version 1.0
// @Accept multipart/form-data
// @Success 100 object util.Result 成功
// @Failure 103/104 object util.Result 失败
// @Router /api/v1/article_files [post]
func ImportArticlesFromFiles(ctx *gin.Context) {
	multiForm, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	files := multiForm.File["file[]"]
	if len(files) > 10 || len(files) == 0 {
		ctx.JSON(http.StatusOK, util.Result{
			Code: util.RequestError,
			Msg:  "上传文件个数不能少于 1 个，也不能多于 10 个",
			Data: nil,
		})
		return
	}
	for _, file := range files {
		// 校验文件类型
		if path.Ext(file.Filename) != ".md" {
			ctx.JSON(http.StatusOK, util.Result{
				Code: util.RequestError,
				Msg:  "只支持导入 md 格式的文件",
				Data: nil,
			})
			return
		}
		// 校验文件大小
		if file.Size > 2*1024*1024 {
			ctx.JSON(http.StatusOK, util.Result{
				Code: util.RequestError,
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
			log.Println("error: ", err.Error())
			ctx.JSON(http.StatusOK, util.Result{
				Code: util.ServerError,
				Msg:  "读取文件内容失败",
				Data: nil,
			})
			return
		}
		// 关闭文件
		src.Close()
		// 读取文件
		bytes, err := ioutil.ReadAll(src)
		article := model.Article{
			Content: string(bytes),
			Title:   util.GetFileNameOnly(file.Filename),
		}
		err = article.SaveFromFile()
		if err != nil {
			log.Println("error: ", err.Error())
			ctx.JSON(http.StatusOK, util.Result{
				Code: util.ServerError,
				Msg:  "数据库错误",
				Data: nil,
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, util.Result{
		Code: util.Success,
		Msg:  "导入成功",
		Data: nil,
	})
}
