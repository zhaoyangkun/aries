package middlewares

import (
	"aries/config/setting"
	"aries/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

// 日志中间件
func LoggerMiddleWare() gin.HandlerFunc {
	logger := initLogger()
	return func(ctx *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		ctx.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := ctx.Request.Method
		// 请求路由
		reqUri := ctx.Request.RequestURI
		// 状态码
		statusCode := ctx.Writer.Status()
		// 请求IP
		clientIP := ctx.ClientIP()
		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime,
			clientIP, reqMethod, reqUri)
		// 控制台打印日志
		logrus.Infof("| %3d | %13v | %s | %s | %s |", statusCode, latencyTime,
			clientIP, reqMethod, reqUri)
	}
}

func initLogger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if rootPath, err := os.Getwd(); err == nil {
		logFilePath = filepath.Join(rootPath, "logs")
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println("err: ", err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := filepath.Join(logFilePath, logFileName)
	// 创建日志文件
	if !utils.FileIsExists(fileName) {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println("err: ", err.Error())
		}
	}
	// 打开文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}

	// 实例化
	logger := logrus.New()
	// 设置输出
	logger.Out = src
	// 设置日志级别
	if setting.Config.Server.Mode == gin.DebugMode {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return logger
}
