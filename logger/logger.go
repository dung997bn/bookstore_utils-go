package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type bookstoreLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

//log type
type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutput()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func getOutput() string {
	output := strings.ToLower(os.Getenv(envLogOutput))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevel() zapcore.Level {
	switch os.Getenv(envLogLevel) {
	case "info":
		return zap.InfoLevel

	case "error":
		return zap.ErrorLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}

//GetLogger func
func GetLogger() bookstoreLogger {
	return log
}

//Info func
func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

//Error func
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}
