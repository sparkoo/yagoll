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
	_, file, line, _ := runtime.Caller(2)
	message.file = file
	message.line = line

	finalMessage := prefix(message)
	if message.format {
		finalMessage += fmt.Sprintf(message.message, message.args...)
	} else {
		finalMessage += message.message
		finalMessage += fmt.Sprint(message.args[0])
		finalMessage += fmt.Sprint(message.args[1:])
	}
	if message.newLine {
		finalMessage += "\n"
	}
	_, _ = err.Write([]byte(fmt.Sprintf("%-50s", finalMessage)))
}

func prefix(message message) string {
	filePath := strings.Split(message.file, "/")
	fileName := filePath[len(filePath)-1]
	prefix := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05.000"), " ", fileName, ":", message.line)
	return fmt.Sprintf("[ %-45s ]> ", prefix)
}

func Println(msg ... interface{}) {
	writeMessage(message{args: msg, newLine: true, level: TRACE})
}

func Printf(msg string, args ... interface{}) {
	writeMessage(message{message: msg, args: args, format: true, level: TRACE})
}

func Printfln(msg string, args ... interface{}) {
	writeMessage(message{message: msg, args: args, format: true, newLine: true, level: TRACE})
}
