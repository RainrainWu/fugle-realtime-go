package logger

import "go.uber.org/zap"

var (
	PrintLogger *zap.SugaredLogger
)

func init() {
	printLoggerProd, _ := zap.NewProduction()
	PrintLogger = printLoggerProd.Sugar()
	defer PrintLogger.Sync()
}
