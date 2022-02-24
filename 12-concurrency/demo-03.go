package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

var opCount int

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
	//opCount++
	mutex.Lock()
	opCount = opCount + 1
	mutex.Unlock()
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
