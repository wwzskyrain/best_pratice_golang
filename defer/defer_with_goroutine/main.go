package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("main 方法开始")
	logWithMethodName("main", "start--->")
	testWhenDeferRunWithGoroutine()
	time.Sleep(10 * time.Second)
	logWithMethodName("main", "end!!!")
}

func testWhenDeferRunWithGoroutine() {
	methodName := "testWhenDeferRunWithGoroutine"
	logWithMethodName(methodName, "start")
	defer func() {
		log.Println("这是defer函数.")
	}()
	go func() {
		log.Println("这是goroutine函数=睡一会")
		time.Sleep(5 * time.Second)
		log.Println("这是goroutine函数=睡好了，结束！")
	}()
	logWithMethodName(methodName, "end")
}

func logWithMethodName(methodName string, logMessage string) {
	// fmt.Printf("[%s] %s\n", methodName, logMessage)
	log.Default().Printf("[%s] %s", methodName, logMessage)
}
