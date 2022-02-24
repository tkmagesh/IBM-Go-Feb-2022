package main

import (
	"fmt"
	"time"
)

/* Keep generating the even number until the user hits ENTER key */

func main() {
	done := make(chan bool)
	evenCh := genEvenNos(done)
	for i := 0; i < 10; i++ {
		no := <-evenCh
		fmt.Println(no)
	}
	done <- true
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
