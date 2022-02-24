package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	for i := 0; i < 5; i++ {
		wg.Add(1) //initialize the counter
		go f1()
	}
	f2()
	wg.Wait() // block until the counter becomes zero
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
