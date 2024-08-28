package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// hmap 结构
type hmap struct {
	count      int
	flags      uint8
	B          uint8
	noverflow  uint16
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra      unsafe.Pointer
}

// 通过反射和 unsafe 获取底层结构
func inspectMap(m interface{}) {
	// 通过反射获取 map 的指针
	val := reflect.ValueOf(m)
	mapValue := (*hmap)(unsafe.Pointer(val.Pointer()))

	// 打印 hmap 结构的各个字段
	fmt.Printf("Count: %d\n", mapValue.count)
	fmt.Printf("Flags: %d\n", mapValue.flags)
	fmt.Printf("B: %d (number of buckets: %d)\n", mapValue.B, 1<<mapValue.B)
	fmt.Printf("Noverflow: %d\n", mapValue.noverflow)
	fmt.Printf("Hash0: %d\n", mapValue.hash0)
	fmt.Printf("Buckets: %v\n", mapValue.buckets)
	fmt.Printf("Oldbuckets: %v\n", mapValue.oldbuckets)
	fmt.Printf("Nevacuate: %d\n", mapValue.nevacuate)
	fmt.Printf("Extra: %v\n", mapValue.extra)
}

func main() {
	aMap := make(map[int]int)
	aMap[1] = 100
	aMap[2] = 200
	inspectMap(aMap)

	var (
		bMap map[int]int
	)
	bMap = make(map[int]int)
	fmt.Printf("current len-> %d\n", len(bMap))
	bMap[1] = 1
	fmt.Printf("current len-> %d\n", len(bMap))

	delete(bMap, 1)
	fmt.Printf("current len-> %d\n", len(bMap))
}
