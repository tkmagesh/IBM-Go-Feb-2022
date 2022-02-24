package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	print("hello")
	print("world")
	fmt.Println("main completed")
}

func print(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
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
