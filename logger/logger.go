/*
logger package
Define logging methods
*/
package logger

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

var (
	traceLogger   *log.Logger // Just about anything
	commandLogger *log.Logger // Important information
	infoLogger    *log.Logger // Important information
	errorLogger   *log.Logger // Critical problem

	// Defining color functions for loggers
	red    = color.New(color.FgRed).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
)

func init() {

	// Initialising loggers
	traceLogger = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Lmicroseconds)

	commandLogger = log.New(color.Output,
		yellow("> "),
		0)

	infoLogger = log.New(color.Output,
		cyan("INFO: "),
		0)

	errorLogger = log.New(color.Output,
		red("ERROR: "),
		0)
}

// Command log executed commands
func Command(message string) {
	commandLogger.Printf("%s\n", yellow(message))
}

// Info iog important information
func Info(message string) {
	infoLogger.Printf("%s\n", message)
}

// Error log error information and exit
func Error(err error) {
	errorLogger.Fatalf("%s\n", err)
}

// Trace log anything to trace
func Trace(message interface{}) {
	traceLogger.Printf("%+v\n", message)
}

// SetTraceOutput change Trace output writer, default is ioutil.Discard.
func SetTraceOutput(w io.Writer) {
	traceLogger.SetOutput(w)
}
