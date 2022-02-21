package main

import "fmt"

func main() {
	var x int32 = 10
	var y int64 = 20
	//result := x + y
	//result := int64(x) + y
	result := x + int32(y)
	fmt.Println(result)
}
