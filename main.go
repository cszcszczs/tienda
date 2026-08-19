package main

import (
	"fmt"
	"strings"
)

type Product struct {
	name  string
	price float64
}

var products = []Product{
	{name: "Mouse", price: 10.0},
	{name: "Teclado", price: 20.0},
	{name: "Audifonos", price: 30.0},
}

func main() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. List products")
		fmt.Println("2. Add product")
		fmt.Println("3. Updated product")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			listProducts()
		case 2:
			var name string
			var price float64
			fmt.Println("Enter product name:")
			fmt.Scanln(&name)
			fmt.Println("Enter product price:")
			fmt.Scanln(&price)
			products = append(products, Product{name: name, price: price})

		case 3:
			listProducts()
			var name string
			var price float64
			fmt.Println("Enter product name:")
			fmt.Scanln(&name)
			fmt.Println("Enter product price:")
			fmt.Scanln(&price)
			for i, product := range products {
				if strings.EqualFold(product.name, strings.TrimSpace(name)) {
					products[i].price = price
					break
				}
			}
		case 4:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func listProducts() {
	for _, product := range products {
		fmt.Printf("Name: %s, Price: %.2f\n", product.name, product.price)
	}
}
