package yagoll

import "log"

func Println(msg ... interface{}) {
	log.Println(msg)
}

func Printf(msg string, args ... interface{}) {
	log.Printf(msg, args)
}
