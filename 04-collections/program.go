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

	//Slice
	fmt.Printf("\nSlice\n")
	//var products []string
	//var products []string = []string{"Pen", "Pencil", "Marker"}
	//var products []string = []string{}

	var products = make([]string, 0, 20)
	products = append(products, "Pen", "Pencil", "Marker")

	/*
		var products = make([]string, 3, 20)
		products[0] = "Pen"
		products[1] = "Pencil"
		products[2] = "Marker"
	*/
	fmt.Printf("%#v, len=%d, capacity=%d\n", products, len(products), cap(products))
	//using the append function
	/*
		fmt.Println("After appending 1 new product")
		products = append(products, "Scribble-Pad")
		fmt.Printf("%#v, len=%d, capacity=%d\n", products, len(products), cap(products))
	*/
	fmt.Println("After appending 4 new products")
	products = append(products, "Stapler", "Mouse", "Adapter", "Charger")
	fmt.Printf("%#v, len=%d, capacity=%d\n", products, len(products), cap(products))

	fmt.Println("After appending 1 more new product")
	products = append(products, "IPad")
	fmt.Printf("%#v, len=%d, capacity=%d\n", products, len(products), cap(products))
}
