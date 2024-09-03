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
	ch1 := make(chan int64, 2)
	ch2 := make(chan int64, 2)
	ch3 := make(chan int64, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			ch1 <- time.Now().Unix()
			time.Sleep(time.Duration(1) * time.Second)

			ch2 <- time.Now().Unix()
			time.Sleep(time.Duration(1) * time.Second)

			ch3 <- time.Now().Unix()
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case t1 := <-ch1:
				fmt.Println("Received from ch1, ", t1)
			case t2 := <-ch2:
				fmt.Println("Received from ch2, ", t2)
			case t3 := <-ch3:
				fmt.Println("Received from ch3, ", t3)
			}
		}
	}()
	wg.Wait()
}
