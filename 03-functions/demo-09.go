package main

import "fmt"

var counter int

func main() {
	fmt.Println(increment())
}

func incrment() int {
	counter += 1
	return counter
}
