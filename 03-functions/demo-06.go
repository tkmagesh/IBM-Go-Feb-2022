/* functions as arguments to other functions */
package main

import "fmt"

type Operation func(int, int) int

func main() {

	/*
		fmt.Println("invocation started")
		fmt.Println(add(100, 200))
		fmt.Println("invocation completed")
	*/
	logOperation(add, 100, 200)
	/*
		fmt.Println("invocation started")
		fmt.Println(subtract(100, 200))
		fmt.Println("invocation completed")
	*/
	logOperation(subtract, 100, 200)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func logOperation(oper Operation, x, y int) {
	fmt.Println("invocation started")
	fmt.Println(oper(x, y))
	fmt.Println("invocation completed")
}
