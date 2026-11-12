package main

import (
	"fmt"
	"time"
)

func run() {
	time.Sleep(2 * time.Second)
	fmt.Println("run1")
}

func run2() {
	time.Sleep(2 * time.Second)
	fmt.Println("run2")
}

func run3() {
	time.Sleep(2 * time.Second)
	fmt.Println("run3")
}

func add(x int, y int, ch chan int) {
	//fmt.Println(x + y)
	ch <- x + y
}

func main() {
	go run() //will run this function in a new thread
	go run2()
	go run3()
	time.Sleep(3 * time.Second) //provides enough time for all go routines to finish running

	//awaiting for a return value in go routines using a channel
	// a channel provides a way of implementing blocking code
	ch := make(chan int)
	// go add(5, 10, ch)
	// x := <-ch //blocking code that first waits a value from the channel so that it can proceed to the next line
	// go add(5, 0, ch)
	// x = <-ch //awaiting for multiple go routings

	//using multiple channels
	ch2 := make(chan int)
	go add(20, 15, ch2)
	go add(10, 15, ch)

	//awaiting values from both channels
	for i := 0; i < 2; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case y := <-ch2:
			fmt.Println(y)
		}
	}

	//fmt.Println(x, y)
}
