package main

import (
	"log"
	"os"
)

type Logger struct {
	logger  *log.Logger
	verbose bool
}

func newLogger(verbose bool) *Logger {
	return &Logger{
		logger:  log.New(os.Stdout, "", 0),
		verbose: verbose,
	}
}
func (l *Logger) Info(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l.verbose {
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.Error(format, v...)
	os.Exit(1)
}
