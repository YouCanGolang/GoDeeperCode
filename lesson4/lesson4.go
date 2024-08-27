package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("切片a长度: %d, 容量: %d\n", len(a), cap(a))
	a = append(a, 6)
	fmt.Printf("切片a长度: %d, 容量: %d\n", len(a), cap(a))

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	arrayPtr := hdr.Data
	array := (*[10]int)(unsafe.Pointer(arrayPtr))
	fmt.Println("底层数组: ", array)

	b := [5]int{1, 2, 3, 4, 5}
	bs := b[1:4]
	fmt.Printf("切片bs index 0: %d, 切片bs长度: %d, 容量: %d\n", bs[0], len(bs), cap(bs))

	c := []int{1, 2, 3, 4, 5}
	fmt.Printf("切片c长度: %d, 容量: %d\n", len(c), cap(c))
	c = append(c, 6)
	fmt.Printf("切片c长度: %d, 容量: %d\n", len(c), cap(c))
	c = append(c, 7)
	fmt.Printf("切片c长度: %d, 容量: %d\n", len(c), cap(c))
	c = append(c, 8, 9, 10, 11)
	fmt.Printf("切片c长度: %d, 容量: %d\n", len(c), cap(c))

	s := make([]int, 0, 5)
	capOld := cap(s)

	for i := 0; i < 2000; i++ {
		s = append(s, i)
		capNew := cap(s)
		if capNew != capOld {
			fmt.Printf("扩容: 旧容量=%d, 新容量=%d\n", capOld, capNew)
			capOld = capNew
		}
	}
}
