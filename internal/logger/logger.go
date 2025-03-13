package logger

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func init() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	Logger = zapLogger.Sugar()
}
