package main

import (
	"errors"
	"fmt"
)

var InvalidChoice = errors.New("invalid choice")

func main() {
	var result int
	for {
		userChoice, err := getUserChoice()
		if err == InvalidChoice {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		result, err = performOperation(userChoice)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Result = ", result)
	}
}

func performOperation(userChoice int) (result int, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.New("Cannot perform the operation")
		}
	}()
	x, y := getOperands()
	switch userChoice {
	case 1:
		result = add(x, y)
	case 2:
		result = subtract(x, y)
	case 3:
		result = multiply(x, y)
	case 4:
		result = divide(x, y)
	}
	return
}

func getOperands() (x int, y int) {
	fmt.Println("Enter 2 nos : ")
	fmt.Scanf("%d %d\n", &x, &y)
	return
}

func getUserChoice() (userChoice int, err error) {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter the choice:")
	fmt.Scanln(&userChoice)
	if userChoice < 1 || userChoice > 5 {
		err = InvalidChoice
	}
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) (result int) {
	result = x / y
	return
}
