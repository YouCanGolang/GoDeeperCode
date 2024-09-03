package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	noBufChan := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("enter goroutine receive 1")
		for ch := range noBufChan {
			fmt.Printf("noBufChan receive 1 -> %d\n", ch)
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("enter goroutine receive 2")
		for ch := range noBufChan {
			fmt.Printf("noBufChan receive 2 -> %d\n", ch)
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("enter goroutine send 1")
		time.Sleep(time.Duration(3) * time.Second)
		for i := 0; i < 100; i++ {
			noBufChan <- i
			fmt.Printf("noBufChan send -> %d\n", i)
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()
	close(noBufChan)

	wg.Wait()
}
