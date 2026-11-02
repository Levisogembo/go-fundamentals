package main

import (
	"fmt"
)

func try(t string) string {
	fmt.Println(t)
	return t
}

func main() {
	// x := "1234"
	// y, err := strconv.Atoi(x)
	// t, err := strconv.ParseInt(x, 2, 2)
	// x := uint8(25)
	// y := 10000
	// z := int(x) < y
	// fmt.Println(z)

	a := 8
	// switch {
	// case a > 1:
	// 	fmt.Println("one")
	// 	fallthrough
	// default:
	// 	fmt.Printf("no typs\t")
	// }
	fmt.Println(try("Golang"), a)

}
