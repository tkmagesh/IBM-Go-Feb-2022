package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type OpCount struct {
	sync.Mutex
	Count int
}

func (opCount *OpCount) Increment() {
	opCount.Lock()
	{
		opCount.Count = opCount.Count + 1
	}
	opCount.Unlock()
}

var opCount OpCount

func main() {
	fmt.Println("main started")
	for i := 0; i < 100; i++ {
		wg.Add(1) //initialize the counter
		go f1(i)
	}
	f2()
	wg.Wait() // block until the counter becomes zero
	fmt.Printf("opCount = %d\n", opCount.Count)
	fmt.Println("main completed")
}

func f1(id int) {
	//opCount++
	opCount.Increment()
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
