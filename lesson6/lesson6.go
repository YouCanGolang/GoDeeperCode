package main

import (
	"fmt"
)

type Animal interface {
	Speak()
}

type Dog struct {
	Name  string
	Voice string
}

func (d *Dog) Speak() {
	fmt.Printf("Name is %s, Speak-> %s\n", d.Name, d.Voice)
}

func Print[T any](text T) {
	fmt.Printf("Print: %v\n", text)
}

func main() {
	var a int
	a = 10
	fmt.Printf("a: %d\n", a)

	var b float64 = 5.5
	c := float64(a) + b
	fmt.Printf("c: %0.2f\n", c)

	var d interface{}
	d = 100
	switch d.(type) {
	case int:
		fmt.Printf("d: int %d\n", d.(int))
	case string:
		fmt.Printf("d: string %d\n", d.(string))
	default:
		fmt.Printf("unknow")
	}

	var animalDog Animal = &Dog{
		Name:  "Dog",
		Voice: "Wang...",
	}
	animalDog.Speak()

	Print(1)
	Print("hello...")
	Print(1.0989)

	type newInt = int
	var e newInt
	e = 1
	fmt.Printf("e: %d\n", e)
}
