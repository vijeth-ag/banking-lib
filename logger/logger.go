package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config.EncoderConfig = encoderConfig

	Log, err = config.Build(zap.AddCallerSkip(1))

	// Log, err = zap.NewProduction(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

}

func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}
