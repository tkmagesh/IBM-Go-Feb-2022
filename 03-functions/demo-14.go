//panic & recovery
package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic - exiting - err=", err)
			return
		}
		fmt.Println("successful exit")
	}()
	if result, err := divide(100, 0); err == DivideByZeroError {
		fmt.Println("Cannot divide by zero")
	} else if err != nil {
		fmt.Println("Something went wrong ", err)
	} else {
		fmt.Println("Result =", result)
	}
}

func divide(x, y int) (result int, err error) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	result = x / y
	return
}
