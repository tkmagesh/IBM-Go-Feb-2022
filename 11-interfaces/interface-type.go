package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = 1.03
	x = true
	x = "this is a string"
	fmt.Println(x)

	var val interface{}
	val = "100"
	if v, ok := val.(int); ok {
		y := v + 200
		fmt.Println(y)
	} else {
		fmt.Println("val is not an int")
	}

	//val = 1000
	//val = "this is a string value"
	//val = true
	//val = []int{3, 1, 4, 2, 5}
	val = 10.004
	switch z := val.(type) {
	case int:
		fmt.Println("val is an int, val + 100 = ", z+100)
	case string:
		fmt.Println("val is a string, len(val) = ", len(z))
	case bool:
		fmt.Println("val is a bool", z)
	case []int:
		fmt.Println("val is a []int with values ", z)
	default:
		fmt.Println("Unknown type")
	}
	fmt.Printf("typeof val = %T\n", val)

}
