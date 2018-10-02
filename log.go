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

type message struct {
	message string
	args    []interface{}
	newLine bool
	format  bool
	file    string
	line    int
}

func writeMessage(message message) {
	finalMessage := prefix(message)
	if message.format {
		finalMessage += fmt.Sprintf(message.message, message.args)
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
	prefix := fmt.Sprint("[ ", time.Now().Format("2006-01-02 15:04:05.000"), " ", fileName, ":", message.line, " ] ")
	return fmt.Sprintf("%-50s", prefix)
}

func Println(msg ... interface{}) {
	_, file, line, _ := runtime.Caller(1)
	writeMessage(message{args: msg, newLine: true, file: file, line: line})
}

func Printf(msg string, args ... interface{}) {
	_, file, line, _ := runtime.Caller(1)
	writeMessage(message{message: msg, args: args, format: true, file: file, line: line})
}

func Printfln(msg string, args ... interface{}) {
	_, file, line, _ := runtime.Caller(1)
	writeMessage(message{message: msg, args: args, format: true, file: file, line: line, newLine: true})
}
