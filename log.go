package unilog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cfg = zap.Config{
	Encoding:         "console",
	Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,

		StacktraceKey: "stacktrace",
	},
}

// newLogger creates a new instance of zap.Logger and returns a pointer to it
func newLogger(logPath string) (*zap.Logger, error) {
	if logPath != "" {
		f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		f.Close()
		cfg.OutputPaths = append(cfg.OutputPaths, logPath)
		cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, logPath)
	}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
