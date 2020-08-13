package middlewares

import (
	"aries/config/setting"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

// 日志中间件
func LoggerMiddleWare() gin.HandlerFunc {
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
		setting.Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()

		// 控制台打印日志
		logrus.Infof("| %3d | %13v | %s | %s | %s |", statusCode, latencyTime,
			clientIP, reqMethod, reqUri)
	}
}

// 初始化 Logger
func InitLogger() {
	logDir := setting.Config.Logger.LogDir
	logName := setting.Config.Logger.LogName
	logLevel := setting.Config.Logger.Level
	logMaxAge := time.Duration(setting.Config.Logger.MaxAge)
	logFormatter := setting.Config.Logger.Formatter

	if err := os.MkdirAll(logDir, 0777); err != nil {
		logrus.Error("err: ", err.Error())
	}
	logFilePath := filepath.Join(logDir, logName)
	logrus.Info("logFilePath: ", logFilePath)

	// 打开文件
	src, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Error("err: ", err.Error())
	}

	// 实例化
	setting.Logger = logrus.New()
	// 设置输出
	setting.Logger.Out = src

	// 设置日志级别
	switch logLevel {
	case "debug":
		setting.Logger.SetLevel(logrus.DebugLevel)
	case "info":
		setting.Logger.SetLevel(logrus.InfoLevel)
	case "warning":
		setting.Logger.SetLevel(logrus.WarnLevel)
	case "error":
		setting.Logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		setting.Logger.SetLevel(logrus.FatalLevel)
	case "panic":
		setting.Logger.SetLevel(logrus.PanicLevel)
	default:
		setting.Logger.SetLevel(logrus.InfoLevel)
	}
	//setting.Logger.SetReportCaller(true)

	// 设置 rotateLogs
	logWriter, err := rotatelogs.New(
		logFilePath+".%Y-%m-%d.log",
		rotatelogs.WithLinkName(logFilePath),
		rotatelogs.WithMaxAge(logMaxAge*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.Errorln("failed to create logs: ", err.Error())
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := &lfshook.LfsHook{}
	// 设置日志格式
	switch logFormatter {
	case "text":
		lfHook = lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	case "json":
		lfHook = lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		lfHook = lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	setting.Logger.AddHook(lfHook)
}
