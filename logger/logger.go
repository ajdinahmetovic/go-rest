package logger

import (
	"time"

	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

//InitLogger func
func InitLogger() {
	sugarLogger = zap.NewExample().Sugar()
}

//Info func for logger
func Info(message string, params ...interface{}) {
	params = append(params, "time", time.Now())

	defer sugarLogger.Sync()
	sugarLogger.Infow(
		message,
		params...,
	)
}

//Error func
func Error(message string, params ...interface{}) {
	params = append(params, "time", time.Now())

	defer sugarLogger.Sync()
	sugarLogger.Errorw(
		message, params...,
	)
}
