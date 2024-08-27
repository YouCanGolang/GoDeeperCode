package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		a int8
		b int16
		c float64
	)

	fmt.Printf("int8  占用 %d 字节 (%d 位)\n", unsafe.Sizeof(a), unsafe.Sizeof(a)*8)
	fmt.Printf("int32 占用 %d 字节 (%d 位)\n", unsafe.Sizeof(b), unsafe.Sizeof(b)*8)
	fmt.Printf("float64 占用 %d 字节 (%d 位)\n", unsafe.Sizeof(c), unsafe.Sizeof(c)*8)

	var d int8 = 10
	printBinaryInt8(d)

	var e int8 = -10
	printBinaryInt8(e)

	var f string = "123456"
	fmt.Printf("原字符串: %s\n", f)

	fb := []byte(f)
	fb[5] = '7'
	fmt.Printf("修改后字符串: %s\n", string(fb))

	fr := []rune(f)
	fr[5] = '7'
	fmt.Printf("修改后字符串: %s\n", string(fr))

	var g string = "字符串探讨"
	fmt.Printf("原字符串: %s\n", g)

	gb := []byte(g)
	gb[4] = '7'
	fmt.Printf("修改后字符串: %s, 长度-> %d\n", string(gb), len(gb))

	gr := []rune(g)
	gr[4] = '7'
	fmt.Printf("修改后字符串: %s, 长度-> %d\n", string(gr), len(gr))
}

func printBinaryInt8(data int8) {
	size := unsafe.Sizeof(data)
	ptr := unsafe.Pointer(&data)

	for i := uintptr(0); i < size; i++ {
		byteValue := *(*byte)(unsafe.Pointer(uintptr(ptr) + i))
		fmt.Printf("%08b ", byteValue)
	}
	fmt.Println()
}
