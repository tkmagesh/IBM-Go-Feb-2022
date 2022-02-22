package main

import "fmt"

func main() {
	//Array
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(len(nos), nos)
	fmt.Printf("%#v\n", nos)

	//iterating an array (using indexer)
	fmt.Println("Iterating using index")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	fmt.Println("Iterating using range construct")
	/*
		for idx, val := range nos {
			fmt.Println(idx, val)
		}
	*/
	for _, val := range nos {
		fmt.Println(val)
	}
}
