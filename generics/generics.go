package main

import "fmt"

type Number interface {
	int | float64 | uint
}

func add[T Number](x T, y T) T {
	return x + y
}

// using generics in maps
func getValues[K comparable, V any](mp map[K]V) []V {
	values := []V{}

	for _, value := range mp {
		values = append(values, value)
	}
	return values
}

// using generics in structs
type GenericSlice[T any] []T

func (g GenericSlice[T]) Print() {
	for _, value := range g {
		fmt.Println(value)
	}
}

type GenericStruct[T any] struct {
	values T
}

func main() {
	//generics allows flexible types in a statically typed language
	value := add(1, 3)
	value1 := add(34.5, 66.7)

	mp := map[string]int{"a": 100, "b": 5, "c": 10}
	res := getValues(mp)

	g := GenericSlice[int]{1, 2, 3}
	g.Print()

	st := GenericStruct[string]{values: "t"}
	fmt.Println(value, value1, res, st)
}
