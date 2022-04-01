package logger

import (
	"context"
	"os"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: os.Getenv("ENV") == "dev",
	})
}

// Error exibe detalhes do erro
func Error(err error) {
	loc := getLocation()

	log.WithFields(log.Fields{
		"function": loc.function,
		"file":     loc.file,
		"line":     loc.line,
	}).Error(err)
}

// ErrorContext exibe detalhes do erro com o contexto
func ErrorContext(ctx context.Context, err error) {
	loc := getLocation()

	log.WithContext(ctx).
		WithFields(log.Fields{
			"function": loc.function,
			"file":     loc.file,
			"line":     loc.line,
		}).Error(err)
}

// Info exibe detalhes do log info
func Info(msg string) {
	loc := getLocation()

	log.WithFields(log.Fields{
		"function": loc.function,
		"file":     loc.file,
		"line":     loc.line,
	}).Info(msg)
}

// InfoContext exibe detalhes do log info com o contexto
func InfoContext(ctx context.Context, msg error) {
	loc := getLocation()

	log.WithContext(ctx).
		WithFields(log.Fields{
			"function": loc.function,
			"file":     loc.file,
			"line":     loc.line,
		}).Info(msg)
}

// Fatal exibe detalhes do erro
func Fatal(err error) {
	loc := getLocation()

	log.WithFields(log.Fields{
		"function": loc.function,
		"file":     loc.file,
		"line":     loc.line,
	}).Fatal(err)
}

func getLocation() location {
	pc, file, line, _ := runtime.Caller(2)
	file = getFileName(file)
	fn := getFuncName(pc)
	return location{
		function: fn,
		file:     file,
		line:     line,
	}
}

func getFileName(file string) string {
	return strings.Replace(file, os.Getenv("LOGGER_PATH_REPLACE"), "", 1)
}

func getFuncName(pc uintptr) string {
	fn := runtime.FuncForPC(pc).Name()
	splitted := strings.Split(fn, "/")
	return splitted[len(splitted)-1]
}
