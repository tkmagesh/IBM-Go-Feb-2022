package main

import "fmt"

func main() {
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var addOperation func(int, int) int
	addOperation = func(x, y int) int {
		return x + y
	}
	fmt.Println(addOperation(100, 200))

	var subtractOperation func(int, int) int
	subtractOperation = func(x, y int) int {
		return x - y
	}
	fmt.Println(subtractOperation(100, 200))

	multiplyOperation := func(x, y int) int {
		return x * y
	}
	fmt.Println(multiplyOperation(100, 200))
}
