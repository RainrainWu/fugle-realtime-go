package logger

import (
	"go.uber.org/zap"
)

var (
	PrintLogger *zap.SugaredLogger
)

func syncLogger(logger *zap.SugaredLogger) {
	err := logger.Sync()
	if err != nil {
		PrintLogger.Error(err.Error())
	}
}

func init() {
	printLoggerProd, _ := zap.NewProduction()
	PrintLogger = printLoggerProd.Sugar()
	defer syncLogger(PrintLogger)
}
