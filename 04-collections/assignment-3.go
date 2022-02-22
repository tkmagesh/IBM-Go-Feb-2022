package main

import (
	"errors"
	"fmt"
)

type Operation func(int, int) int

var InvalidChoice = errors.New("invalid choice")

var operations map[int]Operation = map[int]Operation{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	for {
		userChoice := getUserChoice()
		if userChoice == 5 {
			break
		}
		result, err := performOperation(userChoice)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Result = ", result)
	}
}

func performOperation(userChoice int) (result int, err error) {
	if operation, exists := operations[userChoice]; exists {
		x, y := getOperands()
		result = operation(x, y)
		return
	}
	err = InvalidChoice
	return
}

func getOperands() (x int, y int) {
	fmt.Println("Enter 2 nos : ")
	fmt.Scanf("%d %d\n", &x, &y)
	return
}

func getUserChoice() (userChoice int) {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter the choice:")
	fmt.Scanln(&userChoice)
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
