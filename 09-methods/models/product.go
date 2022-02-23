package models

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discount float64) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}
