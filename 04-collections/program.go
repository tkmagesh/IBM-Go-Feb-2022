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

	fmt.Printf("\nSlicing\n")
	/*
		[lo : hi] => from lo to hi-1
		[lo:] => from lo to len-1
		[:hi] => from 0 to hi-1
		[:] => copy of the slice
	*/
	fmt.Println("products[2:5] => ", products[2:5])
	fmt.Println("products[3:] => ", products[3:])
	fmt.Println("products[:3] => ", products[:3])

	var newProducts = []string{"book", "notebook"}
	products = append(products, newProducts...)
	fmt.Println(products)

	//Map
	fmt.Printf("\nMaps\n")
	//var productRanks map[string]int = map[string]int{"Pen": 1, "Pencil": 4}
	var productRanks map[string]int = map[string]int{
		"Pen":    1,
		"Pencil": 4,
	}
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["Pen"] = 1
	*/
	productRanks["Marker"] = 3
	productRanks["Scribble-Pad"] = 2
	productRanks["Mouse"] = 5
	fmt.Println(productRanks, len(productRanks))

	//iterating a map
	fmt.Println("Iterating a map (using range construct)")
	for key, value := range productRanks {
		fmt.Printf("Rank of %s is %d\n", key, value)
	}

	//checking if a key exists
	var keyToSearch = "Fountain-Pen"
	if value, exists := productRanks[keyToSearch]; exists {
		fmt.Printf("%q exists with value %d\n", keyToSearch, value)
	} else {
		fmt.Printf("%q does not exist\n", keyToSearch)
	}

	//deleting a key/value pair
	delete(productRanks, "Pen")
	fmt.Println("After deleting Pen")
	fmt.Println(productRanks)
}
