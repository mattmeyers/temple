package main

import (
	"log"
	"os"
)

type logger struct {
	logger  *log.Logger
	verbose bool
}

func newLogger(verbose bool) *logger {
	return &logger{
		logger:  log.New(os.Stdout, "", 0),
		verbose: verbose,
	}
}

func (l *logger) Info(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *logger) Debug(format string, v ...interface{}) {
	if l.verbose {
		l.logger.Printf(format, v...)
	}
}

func (l *logger) Error(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *logger) Fatal(format string, v ...interface{}) {
	l.Error(format, v...)
	os.Exit(1)
}
