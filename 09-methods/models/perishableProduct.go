package models

import "fmt"

type PerishableProduct struct {
	Product
	Expiry string
}

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %s", pp.Product.Format(), pp.Expiry)
}
