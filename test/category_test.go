package test

import (
	"aries/config/app"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var TestRouter *gin.Engine

// 初始化路由
func init() {
	TestRouter = app.InitApp()
}

// 测试获取所有分类数据
func TestGetAllCategories(t *testing.T) {
	w := httptest.NewRecorder()                                               // 创建响应
	req := httptest.NewRequest(http.MethodGet, "/api/v1/all_categories", nil) // 建立请求
	req.Header.Add("content-type", "application/json")                        // 设置请求头
	TestRouter.ServeHTTP(w, req)                                              // 发送请求
	assert.Equal(t, w.Code, http.StatusOK)                                    // 校验状态码
	log.Println(w.Body)                                                       // 打印响应体内容
}

// 测试分页获取分类数据
func TestGetCategoriesByPage(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet,
		"/api/v1/categories?page=1&size=2&key=", nil)
	req.Header.Add("content-type", "application/json")
	TestRouter.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	log.Println(w.Body)
}
