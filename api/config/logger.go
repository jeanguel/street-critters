package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

type CustomLogger struct {
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

// MainLogger performs logging for tasks on the main thread
var MainLogger CustomLogger

var mainLogFile, workerLogFile *rotatelogs.RotateLogs

// CreateWorkerLogger returns a custom logger that adds the worker
// name to the logger prefix
func CreateWorkerLogger(workerName string) *CustomLogger {
	return &CustomLogger{
		Debug: log.New(workerLogFile, "— "+workerName+" — DEBUG — ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds),
		Info:  log.New(workerLogFile, "— "+workerName+" — INFO — ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds),
		Warn:  log.New(workerLogFile, "— "+workerName+" — WARN — ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds),
		Error: log.New(workerLogFile, "— "+workerName+" — ERROR — ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}
}

// LoggerMiddleware ensures that all requests are
// logged properly for future analysis
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type logResponseWriter struct {
			http.ResponseWriter
			statusCode int
		}

		logEntry := fmt.Sprintf("%s -> %s %s ", r.RemoteAddr, r.Method, r.URL.String())

		lrw := &logResponseWriter{w, http.StatusOK}
		next.ServeHTTP(lrw, r)

		logEntry += fmt.Sprint(lrw.statusCode)
		MainLogger.Info.Println(logEntry)
	})
}

func initializeMainLogger() {
	var err error

	_, fErr := os.Stat(Config.Application.LogDirectory)
	if os.IsNotExist(fErr) {
		err = os.Mkdir(Config.Application.LogDirectory, 0770)
		if err != nil {
			panic("Unable to create log directory: " + err.Error())
		}
	} else if fErr != nil {
		panic("Unexpected log directory error occured: " + err.Error())
	}

	mainLogFile, err = rotatelogs.New(
		fmt.Sprintf("%s.%s.log", filepath.Join(Config.Application.LogDirectory, "api-server"), "%Y-%m-%d"),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("Unable to set application log rotation: " + err.Error())
	}

	workerLogFile, err = rotatelogs.New(
		fmt.Sprintf("%s.%s.log", filepath.Join(Config.Application.LogDirectory, "api-workers"), "%Y-%m-%d"),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("Cannot access workers log file: " + err.Error())
	}

	MainLogger.Debug = log.New(mainLogFile, "— DEBUG — ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds)
	MainLogger.Info = log.New(mainLogFile, "— INFO — ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds)
	MainLogger.Warn = log.New(mainLogFile, "— WARN — ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds)
	MainLogger.Error = log.New(mainLogFile, "— ERROR — ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func closeLogger() {
	mainLogFile.Close()
	workerLogFile.Close()
}
