package main

import "fmt"

func main() {
	//a slice is a view of an array
	//creating a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:3]
	sl[0] = 100 //modifying a slice
	sl = sl[:4] //increasing the capacity of the slice
	fmt.Println(sl, len(sl), cap(sl))
	//fmt.Printf("%v ", arr)

	//creating slice without an underlying array
	slice := []string{"hello", "world"}
	//appending to a slice
	for x := 0; x < 10; x++ {
		slice = append(slice, "time")
		fmt.Println(slice, len(slice), cap(slice))
	}
	//iterating over the slice
	for i, value := range slice {
		fmt.Println(i, value)
	}

	//specifying the size and capacity of a slice
	specifiedSlice := make([]int, 10, 20) //len is 10 and capacity is 20
	fmt.Println(specifiedSlice)
}

// taking slice as an argument
func test(arr []string) {
	arr[0] = "changed slice"
}
