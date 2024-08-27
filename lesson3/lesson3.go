package main

import (
	"fmt"
	"reflect"
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

type Cat struct {
	Name  string
	Voice string
}

func (c *Cat) Speak() {
	fmt.Printf("Name is %s, Speak-> %s\n", c.Name, c.Voice)
}

func main() {
	var animalDog Animal = &Dog{
		Name:  "Dog",
		Voice: "Wang...",
	}
	animalDog.Speak()

	var animalCat Animal = &Cat{
		Name:  "Dog",
		Voice: "Miaow...",
	}
	animalCat.Speak()

	var (
		a interface{}
	)
	a = 1
	fmt.Printf("类型为: %s\n", reflect.TypeOf(a))
	fmt.Printf("值为: %v\n", reflect.ValueOf(a))
}
