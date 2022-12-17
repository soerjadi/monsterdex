package log

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger = NewZeroLogger(&Config{Level: InfoLevel})
)

func NewZeroLogger(cfg *Config) zerolog.Logger {
	var (
		zrLog  zerolog.Logger
		writer io.Writer = os.Stderr
	)

	// file, err := os.OpenFile(
	// 	cfg.LogFilePath+"app.log",
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY,
	// 	0664,
	// )
	// if err != nil {
	// 	return zrLog
	// } else if file != nil {
	// 	writer = file
	// }

	zerolog.TimestampFieldName = "time"
	zerolog.MessageFieldName = "msg"
	zerolog.CallerFieldName = "line"
	zerolog.ErrorFieldName = "err"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	zrLog = zerolog.New(writer).With().Logger()
	zrLog = setLevel(zrLog, cfg.Level)
	return zrLog
}

func setLevel(zrLog zerolog.Logger, level Level) zerolog.Logger {
	switch level {
	case TraceLevel:
		zrLog = zrLog.Level(zerolog.TraceLevel)
	case DebugLevel:
		zrLog = zrLog.Level(zerolog.DebugLevel)
	case InfoLevel:
		zrLog = zrLog.Level(zerolog.InfoLevel)
	case WarnLevel:
		zrLog = zrLog.Level(zerolog.WarnLevel)
	case ErrorLevel:
		zrLog = zrLog.Level(zerolog.ErrorLevel)
	case FatalLevel:
		zrLog = zrLog.Level(zerolog.FatalLevel)
	default:
		zrLog = zrLog.Level(zerolog.InfoLevel)
	}
	return zrLog
}

func Debug(args ...interface{}) {
	logger.Debug().Timestamp().Msg(fmt.Sprintln(args...))
}

func DebugWithFields(msg string, fields KV) {
	logger.Debug().Timestamp().Fields(fields).Msg(msg)
}

func Info(args ...interface{}) {
	logger.Info().Timestamp().Msg(fmt.Sprintln(args...))
}

func InfoWithFields(msg string, fields KV) {
	logger.Info().Timestamp().Fields(fields).Msg(msg)
}

func Warn(args ...interface{}) {
	logger.Warn().Timestamp().Msg(fmt.Sprintln(args...))
}

func WarnWithFields(msg string, fields KV) {
	logger.Warn().Timestamp().Fields(fields).Msg(msg)
}

func Error(args ...interface{}) {
	logger.Error().Timestamp().Msg(fmt.Sprintln(args...))
}

func ErrorWithFields(msg string, fields KV) {
	logger.Error().Timestamp().Fields(fields).Msg(msg)
}

func Fatal(args ...interface{}) {
	logger.Fatal().Timestamp().Msg(fmt.Sprintln())
}

func FatalWithFields(msg string, fields KV) {
	logger.Fatal().Timestamp().Fields(fields).Msg(msg)
}
