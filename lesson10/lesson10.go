package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int
)

func increment() {
	for i := 0; i < 10000; i++ {
		counter++
	}
}

var (
	counterLock int
	mu          sync.Mutex
)

func incrementLock() {
	mu.Lock()
	for i := 0; i < 10000; i++ {
		counterLock++
	}
	mu.Unlock()
}

func incrementChan(ch chan int) {
	for i := 0; i < 10000; i++ {
		ch <- 1
	}
}

var (
	counterAtomic int32
)

func incrementAtomic() {
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&counterAtomic, 1)
	}
}

var (
	sharedVar int32
)

func writer() {
	atomic.StoreInt32(&sharedVar, 42) // 写操作，带有内存屏障
}

func reader() {
	val := atomic.LoadInt32(&sharedVar) // 读操作，带有内存屏障
	fmt.Printf("sharedVar-> %d\n", val)
}

func main() {
	go increment()
	go increment()

	time.Sleep(time.Second)
	fmt.Println("Counter: ", counter)

	go incrementLock()
	go incrementLock()

	time.Sleep(time.Second)
	fmt.Println("CounterLock: ", counterLock)

	var (
		counterChan int
	)
	ch := make(chan int, 3)
	go incrementChan(ch)
	go incrementChan(ch)

	go func() {
		for value := range ch {
			counterChan += value
		}
	}()
	fmt.Println("CounterChan: ", counterLock)

	go incrementAtomic()
	go incrementAtomic()
	time.Sleep(time.Second)
	fmt.Println("CounterAtomic: ", counterAtomic)

	writer()
	reader()
}
