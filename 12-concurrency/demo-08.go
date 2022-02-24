package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	var ch chan int = make(chan int)
	/*
		ch <- 100	// => writing data into the channel
		<- ch 		// => reading data from the channel
	*/
	wg.Add(1)
	go add(100, 200, ch)
	wg.Wait() //blocked until wg.Done() is invoked (depending on line 30)
	result := <-ch
	fmt.Println("Result =", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result //blocking until a read is initiated (depending on line 21)
	wg.Done()
}
