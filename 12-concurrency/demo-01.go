package main

import (
	"fmt"
)

func main() {
	fmt.Println("main started")
	go f1()
	f2()

	/* DO NOT DO THE FOLLOWING */
	//time.Sleep(1 * time.Millisecond)
	/*
		var input string
		fmt.Scanln(&input)
	*/

	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
