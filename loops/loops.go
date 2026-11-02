package main

import (
	"fmt"
)

func main() {
	for idx := 0; idx <= 10; idx++ {
		//fmt.Println(idx)
	}
	a := 8
	//while loop
	for a < 10 {
		fmt.Println("loop")
		a++
	}
	//looping through strings
	str := "eello word"
	//use range syntax which gives access to the index and character itself
	for _, char := range str {
		fmt.Printf("%T ", char)
		//get the length of a string
		fmt.Println(len(str))
	}
	//fmt.Println(str[0], string(str[0]))
}
