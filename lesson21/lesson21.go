package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func doSomething() error {
	_, err := strconv.ParseInt("1.2x", 10, 64)
	if err != nil {
		return fmt.Errorf("parse failed: %w", err)
	}
	return nil
}

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func doSomething1() error {
	return &MyError{Msg: "Something went wrong"}
}

func mightPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("unexpected situation")
}

func fileOperation() error {
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	// 文件操作...
	return nil
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Success")
	}

	if errors.Is(err, strconv.ErrSyntax) {
		fmt.Println("number ErrSyntax")
	}

	var (
		myErr *MyError
	)
	if errors.As(err, &myErr) {
		fmt.Println("Caught MyError:", myErr.Msg)
	}

	doSomething1()

	fmt.Println("Starting")
	mightPanic()
	fmt.Println("This will not run")
}
