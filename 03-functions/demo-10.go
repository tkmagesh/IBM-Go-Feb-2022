//error handling

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {
	fmt.Println("main commenced")
	result, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Cannot divide by zero")
		return
	}
	if err != nil {
		fmt.Println("Something went wrong ", err)
		return
	}
	fmt.Println("Result =", result)
	fmt.Println("main completed")
}

/*
func divide(x, y int) (int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
*/

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	result = x / y
	return
}
