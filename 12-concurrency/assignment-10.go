package main

import (
	"fmt"
	"time"
)

/* Keep generating the even number until the user hits ENTER key */

func main() {
	done := make(chan bool)
	evenCh := genEvenNos(done)

	go func() {
		var input string
		fmt.Scanln(&input)
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
