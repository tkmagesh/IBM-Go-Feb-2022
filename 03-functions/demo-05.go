package main

import "fmt"

//type alias for a function signature
type operation func(int, int) int

func main() {
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	//var addOperation func(int, int) int
	var addOperation operation
	addOperation = func(x, y int) int {
		return x + y
	}
	fmt.Println(addOperation(100, 200))

	//var subtractOperation func(int, int) int
	var subtractOperation operation
	subtractOperation = func(x, y int) int {
		return x - y
	}
	fmt.Println(subtractOperation(100, 200))

	multiplyOperation := func(x, y int) int {
		return x * y
	}
	fmt.Println(multiplyOperation(100, 200))
}
