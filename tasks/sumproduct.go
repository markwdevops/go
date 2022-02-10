package main

import "fmt"

func main() {

	fmt.Println("Sum or Product")
	var operator string
	fmt.Scanln(&operator)

	fmt.Println("Enter a number")
	var n int
	fmt.Scanln(&n)

	for i := 1; i < n; i++ {
		if operator == "sum" || operator == "Sum" {
			var sm int = n + i
			fmt.Println("The sum of", n, "and", i, "is", sm)
		} else if operator == "product" || operator == "Product" {
			var prod int = n * i
			fmt.Println("The product of", n, "and", i, "is", prod)
		} else {
			fmt.Println("Unknown option please retry and enter sum or product.")
			break
		}
	}

}
