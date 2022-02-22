//panic & recovery
package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic - exiting")
			return
		}
		fmt.Println("successful exit")
	}()
	result := divide(100, 0)
	fmt.Println("Result =", result)
}

func divide(x, y int) int {
	return x / y
}
