package main

import "fmt"

func main() {
	//if construct
	/*
		no := 20
		if no%2 == 0 {
			fmt.Println("is even")
		} else {
			fmt.Println("is odd")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is even\n", no)
	} else {
		fmt.Printf("%d is odd\n", no)
	}

	//for construct
	//version 1.0
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum = %d\n", sum)

	//version 2.0
	//sumulating a 'while' construct
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %d\n", numSum)

	//version 3.0
	//as an infinite loop
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Printf("x = %d\n", x)
}
