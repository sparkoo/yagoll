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
var conf = Config{Level: TRACE, ExcludeFiles: []string{}}

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

type Config struct {
	Level        int
	ExcludeFiles []string
}

func SetLevel(level int) {
	conf.Level = level
}

func SetConfig(config Config) {
	conf = config
}

func writeMessage(message message) {
	if message.level < conf.Level {
		return
	}
	_, file, line, _ := runtime.Caller(2)
	message.file = file
	message.line = line

	for _, excludeFile := range conf.ExcludeFiles {
		if strings.Contains(message.file, excludeFile) {
			return
		}
	}

	finalMessage := prefix(message)
	if message.format {
		finalMessage += fmt.Sprintf(message.message, message.args...)
	} else {
		finalMessage += message.message
		if len(message.args) > 0 {
			for _, toPrint := range message.args {
				finalMessage += fmt.Sprint(" ", toPrint)
			}
		}
	}
	if message.newLine {
		finalMessage += "\n"
	}

	formattedMessage := []byte(fmt.Sprintf("%-50s", finalMessage))

	_, _ = err.Write(formattedMessage)
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
	fileNameWithLine = fmt.Sprintf("%-30s", fileNameWithLine)
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
	message.level = DEBUG
	writeMessage(message)
}

func Debugf(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = DEBUG
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

func Fatal(msg ... interface{}) {
	message := createMessage(msg...)
	message.level = ERROR
	writeMessage(message)
	os.Exit(1)
}

func Fatalf(msg string, args ... interface{}) {
	message := createMessagef(msg, args...)
	message.level = ERROR
	writeMessage(message)
	os.Exit(1)
}
