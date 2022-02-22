package main

import "fmt"

func main() {
	increment := getIncrement()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
	fmt.Println(increment()) //=> 5
}

func getIncrement() func() int { // Step:1 - Outer Function
	var counter int           //Step:2 - Variable in the outer function
	increment := func() int { //Step:3 - Inner Function
		counter += 1 //Step:4 - Variable in the outer function is used in the inner function
		return counter
	}
	return increment //Step:5 - extend the lifetime of the inner function beyond the lifetime of the outer function (in this case - returing the inner function)
}
