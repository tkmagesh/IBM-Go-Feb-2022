package main

import "fmt"

func main() {
	fmt.Println("Is 97 prime ? : ", isPrime((97)))
	fmt.Println("Is 62 prime ? : ", isPrime((62)))
	fmt.Println("100 + 200 = ", add(100, 200))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d, remainder = %d\n", quotient, remainder)
	*/

	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d,\n", quotient)
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
