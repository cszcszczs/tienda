package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

var cart []Product

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. List products")
		fmt.Println("2. Add product")
		fmt.Println("3. Update product")
		fmt.Println("4. Add product to cart")
		fmt.Println("5. Buy cart")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		choiceText := readLine()
		choice, err := strconv.Atoi(choiceText)

		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			listProducts()

		case 2:
			fmt.Print("Enter product name: ")
			name := readLine()

			fmt.Print("Enter product price: ")
			priceText := readLine()
			price, err := strconv.ParseFloat(priceText, 64)

			if err != nil {
				fmt.Println("Invalid price")
				continue
			}

			products = append(products, Product{
				name:  name,
				price: price,
			})

			fmt.Println("Product added successfully")

		case 3:
			listProducts()

			fmt.Print("Enter product name: ")
			name := readLine()

			fmt.Print("Enter new product price: ")
			priceText := readLine()
			price, err := strconv.ParseFloat(priceText, 64)

			if err != nil {
				fmt.Println("Invalid price")
				continue
			}

			found := false

			for i := range products {
				if strings.EqualFold(products[i].name, name) {
					products[i].price = price
					found = true
					fmt.Println("Product updated successfully")
					break
				}
			}

			if !found {
				fmt.Println("Product does not exist")
			}

		case 4:
			listProducts()

			fmt.Println("Enter product names separated by commas.")
			fmt.Println("Example: Mouse, Teclado")
			fmt.Print("Products: ")

			input := readLine()
			names := strings.Split(input, ",")

			added := false

			for _, name := range names {
				name = strings.TrimSpace(name)

				found := false

				for _, product := range products {
					if strings.EqualFold(product.name, name) {
						cart = append(cart, product)
						fmt.Printf("%s added to cart\n", product.name)
						found = true
						added = true
						break
					}
				}

				if !found {
					fmt.Printf("Product not found: %s\n", name)
				}
			}

			if added {
				showCart()
			}

		case 5:
			if len(cart) == 0 {
				fmt.Println("The cart is empty")
				continue
			}

			showCart()

			fmt.Print("Confirm purchase? (yes/no): ")
			confirmation := strings.ToLower(readLine())

			if confirmation == "yes" || confirmation == "y" {
				fmt.Println("Purchase completed successfully")
				cart = nil
			} else {
				fmt.Println("Purchase cancelled")
			}

		case 6:
			fmt.Println("Goodbye")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func listProducts() {
	fmt.Println("\nAvailable products:")

	for _, product := range products {
		fmt.Printf("Name: %s, Price: %.2f\n", product.name, product.price)
	}
}

func showCart() {
	fmt.Println("\nCart:")

	total := 0.0

	for _, product := range cart {
		fmt.Printf("Name: %s, Price: %.2f\n", product.name, product.price)
		total += product.price
	}

	fmt.Printf("Total: %.2f\n", total)
}
