package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func main() {
	products := []Product{
		{name: "Mouse", price: 10.0},
		{name: "Teclado", price: 20.0},
		{name: "Audifonos", price: 30.0},
	}

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
			for _, product := range products {
				fmt.Printf("Name: %s, Price: %.2f\n", product.name, product.price)
			}

		case 2:
			var name string
			var price float64
			fmt.Println("Enter product name:")
			fmt.Scanln(&name)
			fmt.Println("Enter product price:")
			fmt.Scanln(&price)
			products = append(products, Product{name: name, price: price})

		case 3:
			var name string
			var price float64
			fmt.Println("Enter product name:")
			fmt.Scanln(&name)
			fmt.Println("Enter product price:")
			fmt.Scanln(&price)
			for i, product := range products {
				if product.name == name {
					products[i].price = price
					break
				}
			}
			fmt.Println("Error ese producto no existe")
		case 4:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
