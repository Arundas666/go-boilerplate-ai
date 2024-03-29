package main

import (
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type ShoppingCart struct {
	products []*Product
}

func (cart *ShoppingCart) AddProduct(product *Product) {
	cart.products = append(cart.products, product)
}

func (cart *ShoppingCart) CalculateTotal() float64 {
	total := 0.0
	for _, product := range cart.products {
		total += product.Price * float64(product.Quantity)
	}
	return total
}

func main() {
	product1 := &Product{ID: 1, Name: "Mobile Phone", Price: 500.0, Quantity: 2}
	product2 := &Product{ID: 2, Name: "Laptop", Price: 1000.0, Quantity: 1}

	cart := &ShoppingCart{}
	cart.AddProduct(product1)
	cart.AddProduct(product2)

	total := cart.CalculateTotal()
	fmt.Println("Total Price:", total)
}