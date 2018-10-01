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

func prefix() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		filePath := strings.Split(file, "/")
		fileName := filePath[len(filePath)-1]
		return fmt.Sprint("[ ", time.Now().Format(time.RFC3339), " ", fileName, ":", line, " ] ")
	} else {
		panic("failed")
	}
}

func Println(msg ... interface{}) {
	err.Write([]byte(fmt.Sprint(prefix(), msg, "\n")))
}

func Printf(msg string, args ... interface{}) {
	err.Write([]byte(fmt.Sprint(prefix(), fmt.Sprintf(msg, args))))
}
