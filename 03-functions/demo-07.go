//functions as return value

package main

import "fmt"

func main() {
	adder := getAdder()
	fmt.Println(adder(100, 200))
	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(200))
}

func getAdder() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}

func getAdderFor(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
