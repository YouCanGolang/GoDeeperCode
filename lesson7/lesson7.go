package main

import "fmt"

const (
	aa int = 100
)

func main() {
	var a int = 1
	fmt.Printf("a: %d\n", a)

	result := add(a, 20)
	fmt.Printf("sum: %d\n", result)
}

func add(x, y int) int {
	sum := x + y
	return sum
}
