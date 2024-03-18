package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

func GetLogger() (*zap.Logger, error) {
	errConf := zap.NewDevelopmentEncoderConfig()
	errConf.EncodeTime = func(time time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(time.Format("15:04:05 2 Jan 2006"))
	}
	errEncoder := zapcore.NewJSONEncoder(errConf)
	errFile, err := os.OpenFile("errFile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	errSync := zapcore.AddSync(errFile)
	errCore := zapcore.NewCore(errEncoder, errSync, zap.ErrorLevel)

	logConf := zap.NewProductionEncoderConfig()
	logConf.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("15:04:05 2 Jan 2006"))
	}

	logEncoder := zapcore.NewConsoleEncoder(logConf)
	logFile, err := os.OpenFile("app-log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err

	}
	logSync := zapcore.AddSync(logFile)
	logCore := zapcore.NewCore(logEncoder, logSync, zap.InfoLevel)
	tee := zapcore.NewTee(errCore, logCore)
	logger := zap.New(tee, zap.AddCaller())
	log.Println("oook")
	return logger, nil
}
