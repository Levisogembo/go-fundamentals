package main

import (
	"errors"
	"fmt"
)

// in Go most errors are found in compile time
// run time errors can be caught if they were not found during compile time which are called panics

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by 0 ")
	}
	return a / b, nil
}

func deferredFunc() {
	fmt.Println("defer")
	r := recover() //used to recover from a runtime error and saves the error
	if r == nil {
		fmt.Println("no error")
	} else {
		fmt.Println(r)
	}

}

func main() {
	defer deferredFunc() //defer ensures that even after a panic occurs, the defer will still be triggered
	//fmt.Println(divide(1, 0))
	//panic("A crash occured") // invoking a panic
	fmt.Println("run")

	//alternative error handling method
	result, err := divide(1, 0)

	if err != nil {
		fmt.Println("error occured")
	} else {
		fmt.Println(result)
	}

}
