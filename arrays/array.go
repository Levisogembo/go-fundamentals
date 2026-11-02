package main

import "fmt"

func main() {
	//explicit array assignment
	//var arr [2]int
	//fmt.Println(arr)

	//implicit assignment
	arr1 := [2]int{1, 2}
	//nested arrays
	nested := [2][2]int{{1, 2}, {3, 4}}
	//initializing an array with dynamic sizes
	dynamic := [...][]string{{"robot", "go", "node"}, {"robot", "go", "node"}}
	//mutating arrays
	dynamic[0] = []string{"eeel"}
	//looping in the array
	for i, value := range dynamic {
		fmt.Println(i, value)
	}
	fmt.Println(arr1, len(nested), dynamic)
}

// function taking an array as an argument
func test(arr *[1]int) {
	arr[0] = 100
}
