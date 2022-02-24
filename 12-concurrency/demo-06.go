package main

import (
	"fmt"
	"sync"
	"time"
)

/* Communicate by sharing memory - DO NOT DO THIS */
var result int
var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200)
	wg.Wait()
	fmt.Println("Result =", result)
	fmt.Println("main completed")
}

func add(x, y int) {
	time.Sleep(3 * time.Second)
	result = x + y
	wg.Done()
}
