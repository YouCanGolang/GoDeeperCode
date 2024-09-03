package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Produced: ", i)
		time.Sleep(time.Duration(1) * time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Println("Consumed: ", item)
	}
}

func main() {
	var (
		wg sync.WaitGroup
	)
	ch := make(chan int, 5)

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
}
