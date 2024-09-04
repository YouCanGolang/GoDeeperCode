package main

import (
	"fmt"
)

func Print[T any](input T) {
	fmt.Println(input)
}

func Print1[int](input int) {
	fmt.Println(input)
}

func Print2[string](input string) {
	fmt.Println(input)
}

func Add[T int | float64](a, b T) T {
	return a + b
}

type Container[T any] struct {
	value T
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)
	item := s.items[n-1]
	s.items = s.items[:n-1]
	return item
}

// Compare 约束 T 必须可比较
func Compare[T comparable](a, b T) bool {
	return a == b
}

// Stringer 定义一个接口
type Stringer interface {
	String() string
}

// PrintString 泛型函数，T 必须实现 Stringer 接口
func PrintString[T Stringer](item T) {
	fmt.Println(item.String())
}

// Person 实现 Stringer 接口的类型
type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

// Number 泛型约束 T 必须是 int 类型的别名
type Number interface {
	~int
}

func Sum[T Number](a, b T) T {
	return a + b
}

type MyInt int // MyInt 是 int 的别名

func main() {
	Print(1)
	Print("hello")

	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.5, 2.3))

	intContainer := Container[int]{value: 42}
	fmt.Println(intContainer.value)

	stringContainer := Container[string]{value: "hello"}
	fmt.Println(stringContainer.value)

	// 创建一个整数栈
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	fmt.Println(intStack.Pop())

	// 创建一个字符串栈
	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println(stringStack.Pop())

	fmt.Println(Compare(1, 2))
	fmt.Println(Compare("Go", "Go"))

	p := Person{Name: "Alice"}
	PrintString(p)

	var a MyInt = 10
	var b MyInt = 20
	fmt.Println(Sum(a, b))
}
