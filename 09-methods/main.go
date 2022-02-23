package main

import (
	"fmt"
	"methods-demo/models"
)

type MyStr string //alias for string

func (ms MyStr) Length() int {
	return len(ms)
}

type RawProduct models.Product

func (rp RawProduct) Raw() string {
	p := models.Product(rp)
	return fmt.Sprintf("%d-%s-%v", p.Id, p.Name, p.Cost)
}

func main() {
	pen := models.Product{
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
	var grapes = models.PerishableProduct{
		Product: models.Product{
			Id:       200,
			Name:     "Grapes",
			Cost:     50,
			Units:    10,
			Category: "Fruits",
		},
		Expiry: "2 Days",
	}

	fmt.Println(grapes.Format())

	var s = MyStr("This is a sample string")
	fmt.Println(s.Length())

	rawPen := RawProduct(pen)
	fmt.Println(rawPen.Raw())
}
