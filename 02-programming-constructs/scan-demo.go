package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter the name :")
	fmt.Scanln(&name)
	fmt.Printf("Hi %s, Have a nice day!\n", name)

	var no1, no2 int
	fmt.Println("Enter 2 numbers")
	fmt.Scanf("%d %d\n", &no1, &no2)
	//fmt.Scan(&no1, &no2)
	fmt.Println(no1 + no2)
}
