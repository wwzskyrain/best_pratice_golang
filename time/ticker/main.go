package main

import (
	"fmt"
	"time"
)

func main() {
	testTicker()
}

func testTicker() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("刚刚已经调用ticker.Stop()")
}
