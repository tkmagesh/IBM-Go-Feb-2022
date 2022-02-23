package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

/*
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}
*/
/* implementing the fmt.Stringer interface */
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

type ProductPredicate func(Product) bool

/* func (products Products) Format() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%s\n", product.Format())
	}
	return result
} */

/* implementing the fmt.Stringer interface */
func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%s\n", product)
	}
	return result
}

func (products Products) IndexOf(p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(p Product) bool {
	return products.IndexOf(p) != -1
}

/*
func (products Products) FilterCostlyProducts() Products {
	var result Products
	for _, product := range products {
		if product.Cost > 1000 {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	var result Products
	for _, product := range products {
		if product.Category == "Stationary" {
			result = append(result, product)
		}
	}
	return result
}
*/

func (products Products) Filter(predicate ProductPredicate) Products {
	var result Products
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func isCostlyProduct(product Product) bool {
	return product.Cost > 1000
}

func isStationaryProduct(product Product) bool {
	return product.Category == "Stationary"
}

func (products Products) All(predicate ProductPredicate) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func (products Products) Any(predicate ProductPredicate) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
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

	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	fmt.Println("Index of stove = ", products.IndexOf(stove))
	fmt.Println("products include stove ? = ", products.Includes(stove))

	fmt.Println("Filter")
	fmt.Println("Costly Products")
	//costlyProducts := products.FilterCostlyProducts()
	costlyProducts := products.Filter(isCostlyProduct)
	//fmt.Println(costlyProducts.Format())
	fmt.Println(costlyProducts)

	fmt.Println("Stationary Products")
	//stationaryProducts := products.FilterStationaryProducts()
	stationaryProducts := products.Filter(isStationaryProduct)
	//fmt.Println(stationaryProducts.Format())
	fmt.Println(stationaryProducts)

	fmt.Println("All")
	fmt.Println("Are all the products costly products ? :", products.All(isCostlyProduct))

	fmt.Println("Any")
	fmt.Println("Are there any costly products ? :", products.Any(isCostlyProduct))
}

/*
Write the following functions
IndexOf => return the index of the given product in the collection
Includes => return true if the given product is present in the collection else return false
Filter => return products that matches the given criteria
	Use cases:
		1. filter all costly products (cost > 1000)
		2. filter all stationary products (category == stationary)
		3. filter all costly stationary products
All => return true if ALL the products matches the given criteria else return false
	Use cases:
		1. Are all the products costly products (cost > 1000) ?
		2. Are all the products stationary products? (category == "Stationary")

Any => return true if ANY of the products matches the given criteria else returns false
	Use cases:
		1. Are there ANY costly product (cost > 1000) ?
		2. Are there ANY stationary products? (category == "Stationary")


Convert those functions into methods
*/
