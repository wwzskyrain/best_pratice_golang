package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	// Test01NoBufferChannel()
	Test02WriteIntoClosedChan()
}

func Test02WriteIntoClosedChan() {
	ch := make(chan int)
	close(ch)
	log.Printf("read from ch, i = %d\n", <-ch)
	log.Printf("read from ch, i = %d\n", <-ch)
	ch <- 19 // panic: send on closed channel
}

// Test01NoBufferChannel 无缓冲通道，写入动作是被阻塞的，直到下游消费者读取了。
// 一共两个变量，一个是chan的capacity，另一个是write动作是单写还是双写
func Test01NoBufferChannel() {
	// var ch = make(chan int) // 无缓冲通道，写入动作是被阻塞的，直到下游消费者读取了。
	var ch = make(chan int, 1) // 无缓冲通道，写入动作是被阻塞的，直到下游消费者读取了。
	// go writeToChV2(ch, rand.Intn(100))
	go writeToCh(ch, rand.Intn(100))
	time.Sleep(2 * time.Second)
	i := <-ch
	log.Printf("TestNoBufferChannel read from ch, i = %d \n", i)
}

func writeToCh(ch chan int, i int) {
	log.Printf("<<<<before write to ch, i = %d \n", i)
	ch <- i
	log.Printf(">>>>after write to ch, i = %d \n", i)
}

func writeToChV2(ch chan int, i int) {
	log.Printf("<<<<第一次写之前, i = %d \n", i)
	ch <- i
	log.Printf(">>>>第一次写完，第二次写之前 , i = %d \n", i)
	ch <- i
	log.Printf(">>>>第二次写完, i = %d \n", i)
}
