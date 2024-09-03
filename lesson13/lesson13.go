package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 获取系统的 CPU 核心数
	fmt.Println("Number of CPUs:", runtime.NumCPU())

	// 获取当前的 GOMAXPROCS
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// 设置 GOMAXPROCS 为 2
	runtime.GOMAXPROCS(2)
	fmt.Println("Current GOMAXPROCS:", runtime.GOMAXPROCS(0))

	runtime.GOMAXPROCS(2) // 设置 P 的数量为 2

	var (
		wg sync.WaitGroup
	)
	wg.Add(3)

	// 顾客 1
	go func() {
		defer wg.Done()
		fmt.Println("顾客 1 正在点餐")
	}()

	// 顾客 2
	go func() {
		defer wg.Done()
		fmt.Println("顾客 2 正在点餐")
	}()

	// 顾客 3
	go func() {
		defer wg.Done()
		fmt.Println("顾客 3 正在点餐")
	}()

	wg.Wait()
	fmt.Println("所有顾客点餐完毕")
}
