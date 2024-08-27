package main

import (
	"fmt"
	"unsafe"
)

type CacheExample struct {
	a int8  // 占用8位，1字节
	b int32 // 占用32位，4字节
	c int16 // 占用16位，2字节
	d int16 // 占用16位，2字节
}

type CacheExample1 struct {
	a int8  // 占用8位，1字节，内存实际占用4字节
	b int32 // 占用32位，4字节，内存实际占用4字节
	c int8  // 占用8位，1字节，内存实际占用2字节
	d int16 // 占用16位，2字节，内存实际占用2字节
}

type CacheExample2 struct {
	a int8  // 占用8位，1字节，内存实际占用1字节
	c int8  // 占用8位，1字节，内存实际占用1字节
	d int16 // 占用16位，2字节，内存实际占用2字节
	b int32 // 占用32位，4字节，内存实际占用4字节
}

type MethodExample struct {
	Score int16
	Age   int16
}

func (m *MethodExample) Print() {
	fmt.Printf("MethodExample score is-> %d\n", m.Score)
}

func (m *MethodExample) Set(score int16) {
	fmt.Printf("MethodExample set, score-> %d\n", score)
	m.Score = score
}

type MethodExample1 struct {
	Score int16
	Age   int16
}

func (m MethodExample1) Print() {
	fmt.Printf("MethodExample1 score is-> %d\n", m.Score)
}

func (m MethodExample1) Set(score int16) {
	fmt.Printf("MethodExample1 set, score-> %d\n", score)
	m.Score = score
}

func main() {
	var (
		example  CacheExample
		example1 CacheExample1
		example2 CacheExample2
	)
	fmt.Printf("结构体 CacheExample 占用的字节数: %d\n", unsafe.Sizeof(example))
	fmt.Printf("结构体 CacheExample1 占用的字节数: %d\n", unsafe.Sizeof(example1))
	fmt.Printf("结构体 CacheExample2 占用的字节数: %d\n", unsafe.Sizeof(example2))

	fmt.Println("Offset of age:", unsafe.Offsetof(MethodExample{}.Age))

	me := &MethodExample{}
	me.Set(12)
	me.Print()

	var (
		me1 MethodExample1
	)
	me1.Set(12)
	me1.Print()
}
