package main

import (
	"fmt"
	"time"
)

/* Generating even numbers for 5 seconds */
/* Simulating time.After() */

func main() {
	done := After(5 * time.Second)
	evenCh := genEvenNos(done)

	for no := range evenCh {
		fmt.Println(no)
	}
	fmt.Println("main completed")
}

func After(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		ch <- time.Now()
	}()
	return ch
}

func genEvenNos(done <-chan time.Time) <-chan int {
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
