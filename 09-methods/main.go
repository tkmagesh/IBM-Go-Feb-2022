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

func main() {
	pen := Product{
		Id:       100,
		Name:     "pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}

	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())

	fmt.Println("Before applying discount")
	fmt.Println(pen.Format())
	fmt.Println("After applying discount")
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	//perishable product
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

	fmt.Println(grapes.Format())
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discount float64) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %s", pp.Product.Format(), pp.Expiry)
}
