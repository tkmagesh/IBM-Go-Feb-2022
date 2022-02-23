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

func (products Products) Sort(attrName string) error {
	var productComparers = map[string]func(int, int) bool{
		"Id": func(i, j int) bool {
			return products[i].Id < products[j].Id
		},
		"Name": func(i, j int) bool {
			return products[i].Name < products[j].Name
		},
		"Cost": func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		},
		"Units": func(i, j int) bool {
			return products[i].Units < products[j].Units
		},
		"Category": func(i, j int) bool {
			return products[i].Category < products[j].Category
		},
	}
	if comparer, exists := productComparers[attrName]; exists {
		sort.Slice(products, comparer)
		return nil
	}
	return fmt.Errorf("Comparer for %s does not exist", attrName)
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
