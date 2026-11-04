package main

import (
	"fmt"
)

type Shape interface {
	getPerimeter() uint
}

type Triangle struct {
	a uint
	b uint
	c uint
}

func (t Triangle) getPerimeter() uint {
	return t.a + t.b + t.c
}

type Square struct {
	width  uint
	length uint
}

func (sq Square) getPerimeter() uint {
	return sq.length * sq.width
}

func main() {
	var s Shape = Triangle{1, 2, 3}
	var t Shape = Square{4, 5}
	fmt.Println(s.getPerimeter(), t.getPerimeter())
}
