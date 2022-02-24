package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var opCount int64

func main() {
	fmt.Println("main started")
	for i := 0; i < 100; i++ {
		wg.Add(1) //initialize the counter
		go f1(i)
	}
	f2()
	wg.Wait() // block until the counter becomes zero
	fmt.Printf("opCount = %d\n", opCount)
	fmt.Println("main completed")
}

func f1(id int) {

	atomic.AddInt64(&opCount, 1)
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
