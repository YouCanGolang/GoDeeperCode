package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

type MyStruct struct {
	Field1 int
	Field2 string
}

// 创建一个对象池
var structPool = sync.Pool{
	New: func() interface{} {
		return &MyStruct{} // 当对象池中没有可用对象时，创建一个新的结构体
	},
}

// 一个示例函数，返回一个结构体
func getStruct() *MyStruct {
	// 从对象池中获取一个结构体
	s := structPool.Get().(*MyStruct)

	// 重置结构体的字段（可选）
	s.Field1 = 0
	s.Field2 = ""

	return s
}

// 使用完结构体后，将其放回对象池
func releaseStruct(s *MyStruct) {
	structPool.Put(s)
}

func main() {
	// 获取结构体对象
	s1 := getStruct()
	s1.Field1 = 10
	s1.Field2 = "Hello"
	fmt.Printf("s1: %+v\n", s1)

	// 使用完后，将结构体放回对象池
	releaseStruct(s1)

	// 再次获取结构体对象
	s2 := getStruct()
	fmt.Printf("s2: %+v\n", s2)

	for i := 0; i < 10; i++ {
		s := getStruct()          // 从对象池获取一个结构体对象
		s.Field1 = i              // 使用结构体对象
		fmt.Printf("s: %+v\n", s) // 打印结构体对象
		releaseStruct(s)          // 使用完毕后将对象放回对象池
	}

	// 设置新的 GOGC 值为 200
	defaultGc := debug.SetGCPercent(200)
	fmt.Printf("Updated GOGC to 200, old default-> %d\n", defaultGc)

	// 禁用自动GC
	debug.SetGCPercent(-1)

	// 分配一些内存
	data := make([]byte, 30*1024*1024) // 分配10 MiB的内存
	fmt.Println("手动触发GC前的内存状态:")
	printMemStats()

	// 手动触发GC
	runtime.GC()

	fmt.Println("手动触发GC后的内存状态:")
	printMemStats()

	// 防止内存被优化掉
	_ = data

	debug.SetGCPercent(100)

	go func() {
		for {
			fmt.Printf("for: %d\n", time.Now().Unix())
			time.Sleep(time.Duration(2) * time.Second)
		}
	}()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("ListenAndServe err: %v", err)
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
