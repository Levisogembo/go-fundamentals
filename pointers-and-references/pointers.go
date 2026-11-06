package main

import "fmt"

func change(x *int) {
	*x = 100
	//fmt.Println(x)
}

// modifying structs using pointers
type Book struct {
	id    int
	title string
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func main() {

	t := "string"
	l := &t
	*l = "int"

	a := 10
	change(&a)

	b := Book{id: 10, title: "Javascript"}
	b.setTitle("Golang")

	//pointers to pointers
	x := 0
	y := &x
	z := &y
	k := &z
	fmt.Println(a, b)
	fmt.Println(x, *y, **z, ***k, t)

}
