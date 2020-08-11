package middlewares

import (
	"aries/config/setting"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
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
		logger.Infof("| %3d | %13v | %s | %s | %s |", statusCode, latencyTime,
			clientIP, reqMethod, reqUri)

		// 控制台打印日志
		logrus.Infof("| %3d | %13v | %s | %s | %s |", statusCode, latencyTime,
			clientIP, reqMethod, reqUri)
	}
}

func initLogger() *logrus.Logger {
	logDir := setting.Config.Logger.LogDir
	if err := os.MkdirAll(logDir, 0777); err != nil {
		logrus.Error("err: ", err.Error())
	}
	logFilePath := path.Join(logDir, "aries.log")
	// 打开文件
	src, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		logrus.Error("err: ", err.Error())
	}
	defer func() {
		// 关闭文件
		err := src.Close()
		if err != nil {
			logrus.Error("err: ", err.Error())
		}
	}()
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

	// 设置 rotateLogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		logFilePath+".%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logFilePath),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.Error("Failed to create logs: %s", err)
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增 Hook
	logger.AddHook(lfHook)

	return logger
}
