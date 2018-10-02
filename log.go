package yagoll

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var out = os.Stdout
var err = os.Stderr
var conf = config{level: TRACE}

type message struct {
	message string
	args    []interface{}
	newLine bool
	format  bool
	file    string
	line    int
	level   int
}

// levels
const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
)

type config struct {
	level int
}

func SetLevel(level int) {
	conf.level = level
}

func writeMessage(message message) {
	if message.level < conf.level {
		return
	}
	_, file, line, _ := runtime.Caller(2)
	message.file = file
	message.line = line

	finalMessage := prefix(message)
	if message.format {
		finalMessage += fmt.Sprintf(message.message, message.args...)
	} else {
		finalMessage += message.message
		if len(message.args) > 0 {
			finalMessage += " "
			finalMessage += fmt.Sprint(message.args...)
		}
	}
	if message.newLine {
		finalMessage += "\n"
	}
	_, _ = err.Write([]byte(fmt.Sprintf("%-50s", finalMessage)))
}

func prefix(message message) string {
	var level string
	switch message.level {
	case TRACE:
		level = "T"
	case DEBUG:
		level = "D"
	case INFO:
		level = "I"
	case WARN:
		level = "W"
	case ERROR:
		level = "E"
	}
	filePath := strings.Split(message.file, "/")
	fileName := filePath[len(filePath)-1]
	fileNameWithLine := fmt.Sprint(fileName, ":", message.line)
	fileNameWithLine = fmt.Sprintf("%-20s", fileNameWithLine)
	logTime := time.Now().Format("2006-01-02 15:04:05.000")
	prefix := fmt.Sprint("[", level, "] ", logTime, " | ", fileNameWithLine, " || ")
	return prefix
}

func createMessage(msg ... interface{}) message {
	return message{message: fmt.Sprint(msg[0]), args: msg[1:], newLine: true}
}

func createMessagef(msg string, args ... interface{}) message {
	return message{message: msg, args: args, format: true, newLine: true}
}

func Println(msg ... interface{}) {
	message := createMessage(msg...)
	message.level = INFO
	writeMessage(message)
}

func Printf(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = INFO
	message.newLine = false
	writeMessage(message)
}

func Printfln(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = INFO
	writeMessage(message)
}

func Trace(msg ... interface{}) {
	message := createMessage(msg...)
	message.level = TRACE
	writeMessage(message)
}

func Tracef(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = TRACE
	writeMessage(message)
}

func Debug(msg ... interface{}) {
	message := createMessage(msg...)
	message.level = INFO
	writeMessage(message)
}

func Debugf(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = INFO
	writeMessage(message)
}

func Info(msg ... interface{}) {
	message := createMessage(msg...)
	message.level = INFO
	writeMessage(message)
}

func Infof(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = INFO
	writeMessage(message)
}

func Warn(msg ... interface{}) {
	message := createMessage(msg...)
	message.level = WARN
	writeMessage(message)
}

func Warnf(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = WARN
	writeMessage(message)
}

func Error(msg ... interface{}) {
	message := createMessage(msg...)
	message.level = ERROR
	writeMessage(message)
}

func Errorf(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = ERROR
	writeMessage(message)
}
