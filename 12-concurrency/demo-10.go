package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	ch := add(100, 200)
	result := <-ch
	fmt.Println("Result =", result)
	fmt.Println("main completed")
}

func add(x, y int) chan int {
	var ch = make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
