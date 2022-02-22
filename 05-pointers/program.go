package main

import "fmt"

func main() {
	var no int = 10
	//addressOfNo := &no
	var noPtr *int = &no
	fmt.Println(noPtr)

	//deferencing (accessing the value at the pointer)
	var val = *noPtr
	fmt.Println(val)

	fmt.Printf("Before incrementing, no = %d\n", no)
	increment(&no)
	fmt.Printf("After incrementing, no = %d\n", no)

	x, y := 10, 20
	fmt.Printf("Before swapping, x = %d, y = %d\n", x, y)
	swap( /*  */ )
	fmt.Printf("After swapping, x = %d, y = %d\n", x, y)
}

func increment(x *int) {
	fmt.Println("Address of x =", x)
	*x = *x + 1
}

func swap( /*  */ ) {
	/* fill in the blanks */
}
