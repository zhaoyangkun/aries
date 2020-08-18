package middlewares

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger 接收gin框架默认的日志
func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		ip := c.ClientIP()
		//query := c.Request.URL.RawQuery
		//userAgent := c.Request.UserAgent()
		//errors := c.Errors.ByType(gin.ErrorTypePrivate).String()

		c.Next()

		end := time.Now()
		cost := end.Sub(start)

		// 输出日志（终端 + 文件）
		//logger.Info(
		//	"",
		//	zap.Int("status", status),
		//	zap.String("method", method),
		//	zap.String("path", path),
		//	zap.String("ip", ip),
		//	zap.Duration("cost", cost),
		//)
		logger.Sugar().Infof("| %3d | %13v | %s | %s | %s |", status, cost, ip, method, path)
	}
}

// Recover recover 掉项目可能出现的 panic，并使用 zap 记录相关日志
func Recover(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

/*// 日志中间件
func Logger() gin.HandlerFunc {
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

	var lfHook *lfshook.LfsHook
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
}*/
