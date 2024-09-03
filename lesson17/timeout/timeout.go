package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int, ch chan<- int) {
	select {
	case <-time.After(4 * time.Second):
		ch <- id * 2
	case <-ctx.Done():
		fmt.Printf("Worker %d canceled\n", id)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan int)
	for i := 1; i <= 5; i++ {
		go worker(ctx, i, ch)
	}

	for i := 1; i <= 5; i++ {
		select {
		case res := <-ch:
			fmt.Println("Received: ", res)
		case <-ctx.Done():
			fmt.Println("Main canceled")
			return
		}
	}
}
