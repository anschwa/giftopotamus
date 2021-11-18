package logger

import (
	"fmt"
	"io"
	"log"
)

const (
	levelDebug   = "DEBU"
	levelInfo    = "INFO"
	levelWarning = "WARN"
	levelError   = "ERRO"
	levelFatal   = "FATA"
	levelPanic   = "PANI"
)

var logger *log.Logger

func Init(w io.Writer) {
	logger = log.New(w, "", log.LUTC|log.LstdFlags)
}

func logln(level string, v ...interface{}) {
	line := fmt.Sprintln(v...)
	logger.Println(level, line)
}

func logf(level string, format string, v ...interface{}) {
	line := fmt.Sprintf(format, v...)
	logger.Println(level, line)
}

func Debug(v ...interface{}) {
	logln(levelDebug, v...)
}

func Debugf(format string, v ...interface{}) {
	logf(levelDebug, format, v...)
}

func Info(v ...interface{}) {
	logln(levelInfo, v...)
}

func Infof(format string, v ...interface{}) {
	logf(levelInfo, format, v...)
}

func Warn(v ...interface{}) {
	logln(levelWarning, v...)
}

func Warnf(format string, v ...interface{}) {
	logf(levelWarning, format, v...)
}

func Error(v ...interface{}) {
	logln(levelError, v...)
}

func Errorf(format string, v ...interface{}) {
	logf(levelError, format, v...)
}

func Fatal(v ...interface{}) {
	line := fmt.Sprintln(v...)
	logger.Fatalln(levelFatal, line)
}

func Fatalf(format string, v ...interface{}) {
	line := fmt.Sprintf(format, v...)
	logger.Fatalln(levelFatal, line)
}
func Panic(v ...interface{}) {
	line := fmt.Sprintln(v...)
	logger.Panicln(levelPanic, line)
}

func Panicf(format string, v ...interface{}) {
	line := fmt.Sprintf(format, v...)
	logger.Panicln(levelPanic, line)
}
