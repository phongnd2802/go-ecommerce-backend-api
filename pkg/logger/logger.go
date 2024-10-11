package logger

import (
	"os"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)


type LoggerZap struct {
	*zap.Logger
}


func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.LogLevel

	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename: config.FileLogName,
		MaxSize: config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge: config.MaxAge,
		Compress: config.Compress,
	}
	core := zapcore.NewCore(
		encoder, 
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), 
		level)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}


func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}