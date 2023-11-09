package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	log, err = config.Build(zap.AddCallerSkip(1))

	//log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}
func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}
func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}
