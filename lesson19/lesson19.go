package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}

type Person struct {
	Name string
	Age  int
}

type Config struct {
	Host string
	Port int
}

func PrintField(obj interface{}, fieldName string) {
	v := reflect.ValueOf(obj)
	fieldVal := v.FieldByName(fieldName)

	if fieldVal.IsValid() {
		fmt.Printf("%s: %v\n", fieldName, fieldVal)
	} else {
		fmt.Printf("Field %s not found\n", fieldName)
	}
}

func main() {
	var x float64 = 3.4

	// 获取变量的类型
	fmt.Println("type:", reflect.TypeOf(x))

	// 获取变量的值
	fmt.Println("value:", reflect.ValueOf(x))

	// 使用反射修改变量的值
	v := reflect.ValueOf(&x).Elem() // Elem() 获取指针指向的值
	v.SetFloat(7.1)
	fmt.Println("new value:", x)

	// call func
	fn := reflect.ValueOf(add)
	args := []reflect.Value{reflect.ValueOf(3), reflect.ValueOf(4)}

	result := fn.Call(args)
	fmt.Println("Add Result: ", result[0].Int())

	// change struct
	pType := reflect.TypeOf(Person{})
	pValue := reflect.New(pType).Elem()

	pValue.FieldByName("Name").SetString("Alice")
	pValue.FieldByName("Age").SetInt(30)

	person := pValue.Interface().(Person)
	fmt.Printf("Person: %+v\n", person)

	// example
	config := Config{"localhost", 8080}

	PrintField(config, "Host")
	PrintField(config, "Port")
	PrintField(config, "Unknown") // 尝试获取不存在的字段
}
