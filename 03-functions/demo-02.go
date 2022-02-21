//variadic functions

package main

import "fmt"

func main() {
	fmt.Println(sum("xyz", 10))
	fmt.Println(sum("xyz", 10, 20))
	fmt.Println(sum("xyz", 10, 20, 30))
	fmt.Println(sum("xyz", 10, 20, 30, 40, 50))
}

func sum(s string, nos ...int /* nos is a slice (array) of int */) int {
	result := 0
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
