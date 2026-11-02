package main

import (
	"fmt"
	"math"
	//"math"
)

// returning a single value
func add(num1 int, num2 int) int {
	return num1 + num2
}

// returning multiple values
func multiple(val1 int) (string, float64) {
	return "The power is", math.Pow(float64(val1), 5)
}

// a function that returns another function
func getFunc(str string) func(string) string {
	return func(str2 string) string {
		return str + str2
	}
}

// accepting dynamic variables like rest operator, variadic function
func sum(nums ...int) int {
	s := 0
	for _, value := range nums {
		s += value
	}
	return s
}

// named return values
func named(num float64) (n float64) {
	n = math.Ceil(num / 2)
	return
}

func main() {
	value := add(1, 2)
	str, val := multiple(2)
	f1 := getFunc("hello")
	f2 := f1("world")
	s := sum(1, 5, 7, 9393, 2)
	fmt.Println(s, f2, named(5))
	fmt.Println(value, str, val)
}
