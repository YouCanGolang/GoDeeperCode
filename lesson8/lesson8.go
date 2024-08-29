package main

import "fmt"

func main() {
	x := createPointer()
	fmt.Printf("x: %d\n", *x)

	y := createClosure()
	fmt.Printf("y: %d\n", y())
}

func createPointer() *int {
	x := 10
	return &x
}

func createClosure() func() int {
	y := 20
	return func() int {
		return y
	}
}
