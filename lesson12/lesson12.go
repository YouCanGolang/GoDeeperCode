package main

import (
	"fmt"
)

const a = 3
const b = 4
const c = a * b

func add(a, b int) int {
	return a + b
}

func main() {
	result := add(2, 3) // 编译器可能会将此调用内联展开
	fmt.Println(result)

	resultO := 2 + 3
	fmt.Println(resultO)

	fmt.Println(c)
}
