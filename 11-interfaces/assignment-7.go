package main

import (
	"fmt"
	"sort"
)

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

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(byName{products})
	case "Cost":
		sort.Sort(byCost{products})
	case "Units":
		sort.Sort(byUnits{products})
	case "Category":
		sort.Sort(byCategory{products})
	default:
		sort.Sort(products)
	}

}

func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

//sort by Name
type byName struct {
	Products
}

func (products byName) Less(i, j int) bool {
	return products.Products[i].Name < products.Products[j].Name
}

//sort by Cost
type byCost struct {
	Products
}

func (products byCost) Less(i, j int) bool {
	return products.Products[i].Cost < products.Products[j].Cost
}

//sort by Units
type byUnits struct {
	Products
}

func (products byUnits) Less(i, j int) bool {
	return products.Products[i].Units < products.Products[j].Units
}

//sort by Name
type byCategory struct {
	Products
}

func (products byCategory) Less(i, j int) bool {
	return products.Products[i].Category < products.Products[j].Category
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

	fmt.Println("Default Sort")
	//sort.Sort(products)
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println("Sort by name")
	//sort.Sort(byName{products})
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println("Sort by cost")
	//sort.Sort(byCost{products})
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println("Sort by units")
	//sort.Sort(byUnits{products})
	products.Sort("Units")
	fmt.Println(products)

	fmt.Println("Sort by category")
	//sort.Sort(byCategory{products})
	products.Sort("Category")
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
