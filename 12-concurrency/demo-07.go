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
	result := <-ch
	wg.Wait()
	fmt.Println("Result =", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}
