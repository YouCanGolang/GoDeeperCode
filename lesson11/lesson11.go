package main

import "fmt"

var (
	globalVar = "I am a global variable"
)

func main() {
	localVar := "I am a local variable"
	fmt.Println(globalVar)
	fmt.Println(localVar)
}
