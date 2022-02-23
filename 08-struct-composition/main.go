package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
}

var x int

func main() {

	var grapes = PerishableProduct{
		Product: Product{
			Id:       200,
			Name:     "Grapes",
			Cost:     50,
			Units:    10,
			Category: "Fruits",
		},
		Expiry: "2 Days",
	}
	//fmt.Println(grapes.Name)
	fmt.Println(grapes.Product.Name)
	fmt.Println(FormatPerishableProduct(grapes))
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s, Expiry = %s", pp.Id, pp.Name, pp.Cost, pp.Units, pp.Category, pp.Expiry)
}
