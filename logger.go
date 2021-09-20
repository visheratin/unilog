package unilog

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLog(cfg zap.Config) {
	var err error
	log, err = cfg.Build()
	if err != nil {
		fmt.Println("unable to initialize log")
		fmt.Println(err)
		log = defaultLog()
	}
}

func SetLogger(l *zap.Logger) {
	log = l
}

func Logger() *zap.Logger {
	if log == nil {
		cfg := DefaultConfig()
		InitLog(cfg)
	}
	return log
}

func defaultLog() *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)
	return zap.New(core)
}
