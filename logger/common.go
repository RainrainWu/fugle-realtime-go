package logger

import (
	"log"

	"go.uber.org/zap"
)

var (
	PrintLogger *zap.SugaredLogger
)

func syncLogger(logger *zap.SugaredLogger) {
	err := logger.Sync()
	if err != nil {
		log.Println(err.Error())
	}
}

func init() {
	printLoggerProd, _ := zap.NewProduction()
	PrintLogger = printLoggerProd.Sugar()
	defer syncLogger(PrintLogger)
}
