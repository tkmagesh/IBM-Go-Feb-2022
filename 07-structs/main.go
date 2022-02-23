package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	//var product Product
	//var product Product = Product{100, "Pen", 10, 100, "Stationary"}
	//var product Product = Product{Id: 100, Name: "Pen", Cost: 10, Units: 100, Category: "Stationary"}
	/*
		var product Product = Product{
			Id:       100,
			Name:     "Pen",
			Category: "Stationary",
		}
	*/
	/*
		product := Product{
			Id:       100,
			Name:     "Pen",
			Cost:     10,
			Category: "Stationary",
		}
	*/
	/*
		product := &Product{
			Id:       100,
			Name:     "Pen",
			Cost:     10,
			Category: "Stationary",
		}
	*/
	product := new(Product) //allocates memory and returns a pointer
	product.Id = 100
	product.Name = "Pen"
	product.Cost = 10
	product.Units = 100
	product.Category = "Stationary"

	//fmt.Println(product)
	fmt.Printf("%#v\n", product)

	//Accessing the attributes/fields
	fmt.Println(product.Id)
	fmt.Println("Before applying discount")
	fmt.Println(Format(*product))
	fmt.Println("After applying discount")
	ApplyDiscount(product, 10)
	fmt.Println(Format(*product))

}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(p *Product, discount float64) {
	p.Cost = p.Cost * ((100 - discount) / 100)
	//(*p).Cost = (*p).Cost * ((100 - discount) / 100)
}
