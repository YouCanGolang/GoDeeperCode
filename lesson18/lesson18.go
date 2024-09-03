package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// 模拟一些 CPU 密集型的工作
	for i := 0; i < 100000000; i++ {
	}
	time.Sleep(time.Duration(3) * time.Second)

	ft, _ := os.Create("trace.out")
	trace.Start(ft)
	defer trace.Stop()

	go func() {
		for i := 0; i < 100000000; i++ {

		}
	}()
	time.Sleep(time.Duration(3) * time.Second)
}
