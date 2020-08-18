package log

import (
	"aries/utils"
	"log"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志
var Logger *zap.Logger

// 初始化日志
func InitLogger(fileName, level, format string, maxSize, maxBackups, maxAge int) (err error) {
	writeSyncer, err := getLogWriter(fileName, maxSize, maxBackups, maxAge)
	if err != nil {
		return
	}

	encoder := getEncoder(format)

	l := new(zapcore.Level)
	err = l.UnmarshalText([]byte(level))
	if err != nil {
		return
	}

	core := zapcore.NewCore(encoder, writeSyncer, l)
	//Logger = zap.New(core, zap.AddCaller())
	Logger = zap.New(core)

	return
}

func getEncoder(format string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	switch format {
	case "json":
		return zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		return zapcore.NewConsoleEncoder(encoderConfig)
	default:
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

func getLogWriter(fileName string, maxSize, maxBackups, maxAge int) (zapcore.WriteSyncer, error) {
	if !utils.FileIsExists(fileName) {
		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		log.Println("日志文件路径：", fileName)
		defer func() {
			_ = file.Close()
		}()
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}

	syncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(lumberJackLogger),
	)

	return syncer, nil
}
