package main

import (
	"fmt"
	"time"
)

func main() {
	evenCh := genEvenNos()
	for no := range evenCh {
		fmt.Println(no)
	}
	fmt.Println("main completed")
}

func genEvenNos() <-chan int {
	var evenNoCh = make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			evenNoCh <- i * 2
		}
		fmt.Println("All even nos are generated")
		close(evenNoCh)
	}()
	return evenNoCh
}
