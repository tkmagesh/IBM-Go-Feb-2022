package main

import "fmt"

var counter int

func main() {
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	counter = 1000
	fmt.Println(increment()) //=> 1001
	fmt.Println(increment()) //=> 1002
}

func increment() int {
	counter += 1
	return counter
}
