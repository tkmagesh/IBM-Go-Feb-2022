package main

import "fmt"

func main() {
	var userChoice, x, y, result int

	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter the choice:")
		fmt.Scanln(&userChoice)
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		fmt.Println("Enter 2 nos : ")
		fmt.Scanf("%d %d\n", &x, &y)
		switch userChoice {
		case 1:
			result = x + y
		case 2:
			result = x - y
		case 3:
			result = x * y
		case 4:
			result = x / y
		}
		fmt.Println("Result = ", result)
	}
}
