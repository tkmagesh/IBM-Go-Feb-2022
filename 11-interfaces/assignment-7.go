package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

/* implementing the fmt.Stringer interface */
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

/* implementing the fmt.Stringer interface */
func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%s\n", product)
	}
	return result
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	//fmt.Println(products.Format())
	fmt.Println(products)

}

/*
	implement sort functionality for the Products
	IMPORTANT : use sort.Sort() function
	Use Cases:
		1. Sort the products by Id
		2. Sort the products by Name
		3. Sort the products by Cost
		4. Sort the products by Units
		5. Sort the products by Category
*/
