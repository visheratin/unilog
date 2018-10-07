package unilog

import "go.uber.org/zap"

var log *zap.Logger

func InitLog(logPath string) error {
	var err error
	log, err = newLogger(logPath)
	return err
}

func Logger() *zap.Logger {
	if log == nil {
		InitLog("")
	}
	return log
}
