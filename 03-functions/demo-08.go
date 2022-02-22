/* functions as arguments to other functions */
package main

import "fmt"

type Operation func(int, int) int

func main() {
	//add(100,200)
	//logOperation(add, 100, 200)
	logAdd := getLoggedOperation(add)
	logAdd(100, 200)

	//subtract(100,200)
	//logOperation(subtract, 100, 200)
	logSubtract := getLoggedOperation(subtract)
	logSubtract(100, 200)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func getLoggedOperation(oper Operation) Operation {
	return func(x, y int) int {
		fmt.Println("invocation started")
		result := oper(x, y)
		fmt.Println("Result = ", result)
		fmt.Println("invocation completed")
		return result
	}
}
