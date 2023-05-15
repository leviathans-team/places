package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang-pkg/config"
	"os"
	"runtime"
)

type Logger interface {
	InitLogger() error
	Debug(msg string)
	Debugf(template string, args ...interface{})
	Info(msg string)
	Infof(template string, args ...interface{})
	Warn(msg string)
	Warnf(template string, args ...interface{})
	Error(err error)
	Errorf(template string, args ...interface{})
	Fatal(msg string)
	Fatalf(template string, args ...interface{})
	Panic(msg string)
	Panicf(template string, args ...interface{})
	Write(p []byte) (n int, err error)
}

type ServiceLogger struct {
	cfg    *config.Config
	logger zerolog.Logger
}

func NewServiceLogger(cfg *config.Config) *ServiceLogger {
	return &ServiceLogger{cfg: cfg}
}

func (l *ServiceLogger) InitLogger() error {
	var w zerolog.LevelWriter
	l.logger = log.With().Caller().Logger()
	w = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout})
	l.logger = zerolog.New(w).Level(zerolog.InfoLevel).With().CallerWithSkipFrameCount(1).Timestamp().Logger().Hook(l)

	return nil
}

func (l *ServiceLogger) Run(e *zerolog.Event, level zerolog.Level, message string) {
	panic("implement me")
}

func (l *ServiceLogger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *ServiceLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *ServiceLogger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *ServiceLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *ServiceLogger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *ServiceLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *ServiceLogger) Error(err error) {
	l.logger.Error().Msg(err.Error())
}

func (l *ServiceLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *ServiceLogger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}

func (l *ServiceLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}

func (l *ServiceLogger) Panic(msg string) {
	l.logger.Panic().Msg(msg)
}

func (l *ServiceLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panic().Msgf(template, args...)
}

func (l *ServiceLogger) Write(p []byte) (n int, err error) {
	// i'll implement this later when we will need to implement our own Fiber logger
	panic("implement me")
}

func (l *ServiceLogger) ErrorFull(error error) {
	_, fn, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("ERROR:\n%s :: %d :: %s", fn, line, error.Error())
	l.logger.Error().Stack().Err(error).Msg(msg)
}
