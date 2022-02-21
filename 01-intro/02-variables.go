package main

import "fmt"

/* package level variables and types*/
//msg := "Hello World" // := NOT allowed at package level

//var msg = "Hello World"

var msg string

func main() {
	/*
		var msg string
		msg = "Hello World"
	*/

	/* var msg string = "Hello World" */

	//type inference
	/* var msg = "Hello World" */

	//msg := "Hello World" /* => This sytax is applicable ONLY in a function (not at package level) */

	msg = "Hello World"
	print()

	//Handling multiple variables
	/*
		var x int
		var y int
		var result int
		var message string

		x = 100
		y = 200
		result = x + y
		message = "Result ="

		fmt.Println(message, result)
	*/

	/*
		var x, y, result int
		var message string

		x = 100
		y = 200
		result = x + y
		message = "Result ="

		fmt.Println(message, result)
	*/

	/*
		var x, y int = 100, 200
		var message string = "Result ="
		var result int = x + y
		fmt.Println(message, result)
	*/

	/*
		var (
			x, y, result int
			message      string
		)

		x = 100
		y = 200
		result = x + y
		message = "Result ="

		fmt.Println(message, result)
	*/

	/*
		var (
			x, y, result int    = 100, 200, 0
			message      string = "Result ="
		)
		result = x + y
		fmt.Println(message, result)
	*/

	/*
		x, y, result, message := 100, 200, 0, "Result ="
		result = x + y
		fmt.Println(message, result)
	*/

	x, y, message := 100, 200, "Result ="
	result := x + y
	fmt.Println(message, result)
}

func print() {
	fmt.Println(msg)
}
