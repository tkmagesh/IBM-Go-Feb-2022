package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string)
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(2)
	go print("hello", ch1, ch2, wg)
	go print("world", ch2, ch1, wg)
	ch1 <- "start"
	wg.Wait()
	fmt.Println("main completed")
}

func print(str string, in chan string, out chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
		out <- "done"
	}
	wg.Done()
}

/*
expected outcome:
main started
hello
world
hello
world
hello
world
hello
world
hello
world
main completed

IMPORTANT:
	The loop should be IN the print function
*/
