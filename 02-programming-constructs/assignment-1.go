package main

import "fmt"

func main() {
	for no := 3; no <= 100; no++ {
		isPrime := true
		for i := 2; i < no; i++ {
			if no%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d is Prime\n", no)
		}
	}
}
