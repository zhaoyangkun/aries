package test

import (
	"aries/config/app"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

// 初始化路由
func init() {
	router = app.InitApp()
}

// 测试获取分类数据功能
func TestGetCategories(t *testing.T) {
	w := httptest.NewRecorder()                                           // 创建响应
	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories", nil) // 建立请求
	req.Header.Add("content-type", "application/json")                    // 设置请求头
	router.ServeHTTP(w, req)                                              // 发送请求
	assert.Equal(t, w.Code, http.StatusOK)                                // 校验状态码
	log.Println(w.Body)                                                   // 打印响应体内容
}
