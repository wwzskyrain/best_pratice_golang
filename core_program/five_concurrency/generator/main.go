package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for i := 0; ; i++ {
			select {
			case <-done:
				// 如果done通道关闭，则跳出循环
				break Label
			case ch <- rand.Int(): // 不担心这个for循环很快的运行吗？不担心，因为这个ch是单空间的。
				fmt.Printf("GenerateIntA ->")
			}
		}
		// 生产者负责关闭通道
		close(ch)
	}()
	return ch
}

func GenerateIntB(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for i := 0; ; i++ {
			select {
			case <-done:
				// 如果done通道关闭，则跳出循环
				break Label
			case ch <- rand.Int():
				fmt.Printf("GenerateIntB ->")
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	doneToAB := make(chan struct{}) // 这个chan就只是一个提线木偶的提线，一拉那边就关闭了
	resultFromA := GenerateIntA(doneToAB)
	resultFromB := GenerateIntB(doneToAB)
	go func() {
	Label:
		for {
			select {
			case ch <- <-resultFromA:
			case ch <- <-resultFromB:
			case <-done:
				// 如果done通道关闭，则跳出循环
				fmt.Printf("GenerateInt done = %+v", done)
				doneToAB <- struct{}{}
				doneToAB <- struct{}{} // 少一个都不行
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	doneChan := make(chan struct{})
	resultChan := GenerateInt(doneChan)
	for i := 0; i < 10; i++ {
		fmt.Printf("No.%d = %d \n", i, <-resultChan)
		time.Sleep(500 * time.Millisecond)
	}
	doneChan <- struct{}{}
	// time.Sleep(time.Second)
	println("main over!")
}
