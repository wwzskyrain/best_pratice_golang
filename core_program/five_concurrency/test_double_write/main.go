package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{})
	ch := GenerateInt(done)
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		log.Printf("No. %d r = %d\n", i, <-ch)
	}
	done <- struct{}{}
	log.Printf("should have last one, r = %d\n", <-ch)
	log.Printf("another one ?  r = %d\n", <-ch)
	println("main done")
}

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	doneToAB := make(chan struct{}) // 这个chan就只是一个提线木偶的提线，一拉那边就关闭了
	resultFromA := GenerateIntA(doneToAB)
	go func() {
	Label:
		for {
			select {
			// case i := <-resultFromA:  //这种写法就有一个死锁，但是没有分析出来
			// 	log.Printf("after A before int")
			// 	ch <- i
			// 	log.Printf("GenerateInt -> ")
			case ch <- <-resultFromA:
			case <-done:
				// 如果done通道关闭，则跳出循环
				fmt.Printf("done")
				doneToAB <- struct{}{}
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case <-done:
				// 如果done通道关闭，则跳出循环
				break Label
			case ch <- rand.Int():
				log.Printf("GenerateIntAAA -> ")
			}
		}
		close(ch)
	}()
	return ch
}
