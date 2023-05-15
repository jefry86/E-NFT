package initialization

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"nft_platform/global"
	"os"
)

func initZap() {
	core := zapcore.NewCore(getEncoder(), getWriter(), getLevel())
	global.Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	//global.Logger = zap.New(core, zap.AddCaller())
	global.SLogger = global.Logger.Sugar()
}

func getLevel() zapcore.Level {
	level := global.Conf.Logger.Level
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getWriter() zapcore.WriteSyncer {
	if global.Conf.Logger.Type != "file" {
		return zapcore.AddSync(os.Stdout)
	}
	path := global.Conf.Logger.Path
	filename := "app.log"
	file := fmt.Sprintf("%s/%s", path, filename)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	if global.Conf.Logger.Type != "file" {
		return zapcore.NewConsoleEncoder(getEncoderConfig())
	} else {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return encoderConfig
}
