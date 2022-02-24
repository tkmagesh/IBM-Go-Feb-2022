package main

import (
	"fmt"
	"time"
)

/* Generating even numbers for 5 seconds */

func main() {
	done := make(chan bool)
	evenCh := genEvenNos(done)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	for no := range evenCh {
		fmt.Println(no)
	}
	fmt.Println("main completed")
}

func genEvenNos(done chan bool) <-chan int {
	var evenNoCh = make(chan int)
	go func() {
		no := 0
		for {
			select {
			case <-done:
				close(evenNoCh)
				return
			case evenNoCh <- no * 2:
				time.Sleep(500 * time.Millisecond)
				no++
			}
		}
	}()
	return evenNoCh
}
