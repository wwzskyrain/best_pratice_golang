package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	go slowTask(ctx)
	// give some time for the slowTask goroutine to conclude
	time.Sleep(3 * time.Second)
	println("main stop")
}
func slowTask(ctx context.Context) {
	select {
	// simulate a time-consuming task using time.After
	case <-time.After(time.Second * 3):
		fmt.Println("Operation complete.")
	case <-ctx.Done():
		// print the cancellation message with the reason using ctx.Err().
		fmt.Println("Canceled:", ctx.Err())
	}
}
