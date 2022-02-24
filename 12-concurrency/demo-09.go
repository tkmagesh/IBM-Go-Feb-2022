package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	var ch chan int = make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("Result =", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result
}
