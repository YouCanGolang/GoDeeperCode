package main

import (
	"fmt"
	"sync"
)

func process(id int, ch chan<- int) {
	defer close(ch)
	result := id * 2
	fmt.Printf("Process %d done\n", id)
	ch <- result
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var (
		wg sync.WaitGroup
	)
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out) // 确保所有 Goroutine 处理完毕后关闭输出通道
	}()

	return out
}

func main() {
	channels := make([]chan int, 3)
	for i := range channels {
		channels[i] = make(chan int)
		go process(i+1, channels[i])
	}

	resultCh := fanIn(channels[0], channels[1], channels[2])
	for result := range resultCh {
		fmt.Println("Received:", result)
	}
}
