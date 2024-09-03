package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	mu      sync.Mutex
	counter int
)

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

var (
	muC   sync.Mutex
	cond  = sync.NewCond(&muC)
	ready = false
	wg    sync.WaitGroup
)

func waitCondition() {
	defer wg.Done()
	cond.L.Lock()
	for !ready {
		cond.Wait()
	}
	// 执行其他操作
	fmt.Println("Cond unlock")
	cond.L.Unlock()
}

func signalCondition() {
	defer wg.Done()
	fmt.Println("signalCondition")
	cond.L.Lock()
	ready = true
	cond.Signal() // 唤醒一个等待的 Goroutine
	cond.L.Unlock()
}

var (
	counterAuto int64
)

func incrementAuto() {
	atomic.AddInt64(&counterAuto, 1)
}

func main() {
	increment()

	wg.Add(2)
	go waitCondition()

	time.Sleep(time.Duration(2) * time.Second)

	go signalCondition()

	wg.Wait()
}
