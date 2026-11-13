package main

import (
	"fmt"
	"sync"
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

func r(ch chan bool) {
	<-ch
}

type Counter struct {
	value int
	lock  sync.Mutex //implementing locks and mutexes to prevent multiple threads from accessing the same memory address and changing it at the same time
}

func count(counter *Counter, wg *sync.WaitGroup) {
	counter.lock.Lock() //acquiring the counter
	defer counter.lock.Unlock()
	counter.value++
	fmt.Println(counter.value)
	//ch <- true
	wg.Done()
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

	//implementing buffer channels
	ch3 := make(chan bool, 2) // creating a buffered channel that can receive 2 different values
	//go r(ch3)
	ch3 <- true
	<-ch3
	fmt.Println("done")

	counter := Counter{}
	//counterChanel := make(chan bool)

	wg := sync.WaitGroup{} //creating a wait group
	wg.Add(100)            //adding 100 go routines to wait for

	for i := 0; i < 100; i++ {
		go count(&counter, &wg)
	}
	//waiting for results to finish
	//1. using channels
	// for i := 0; i < 100; i++ {
	// 	<-counterChanel //await for all go routines to finish
	// }
	//2. using time sleep
	//time.Sleep(2 * time.Second)

	//3. using a wait group
	wg.Wait()
	//fmt.Println(x, y)
}
