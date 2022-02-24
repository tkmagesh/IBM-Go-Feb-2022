package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float32 `json:"cost"`
	Units    int     `json:"-"`
	Category string  `json:"category"`
}

var products = []Product{
	{Id: 100, Name: "Pen", Cost: 10, Units: 100, Category: "Stationary"},
	{Id: 101, Name: "Pencil", Cost: 5, Units: 200, Category: "Stationary"},
	{Id: 102, Name: "Marker", Cost: 50, Units: 50, Category: "Stationary"},
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World!"))
	})
	http.HandleFunc("/products", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			encoder := json.NewEncoder(rw)
			encoder.Encode(products)
		case "POST":
			decoder := json.NewDecoder(r.Body)
			var newProduct Product
			decoder.Decode(&newProduct)
			products = append(products, newProduct)
			rw.WriteHeader(http.StatusCreated)
			encoder := json.NewEncoder(rw)
			encoder.Encode(newProduct)
		case "PUT":
			rw.Write([]byte("The given product is updated"))
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}

	})
	http.ListenAndServe(":8080", nil)
}
