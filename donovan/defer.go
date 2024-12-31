package edu

import (
	"fmt"
	"log"
	"time"
)

func slowOperation() {
	defer trace("slow operation")() // без скобок в конце будет выполнен только "вход"
	fmt.Println("some move")
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("вход в %s", msg)
	return func() { log.Printf("выход из %s (%s)", msg, time.Since(start)) }
}

func Defer() {
	slowOperation()
}
