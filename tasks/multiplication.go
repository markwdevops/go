package main

import "fmt"

func main() {

	fmt.Println("Enter a number for it's multiplication table.")
	var n int
	fmt.Scanln(&n)

	var times = 12

	for i := 0; i <= times; i++ {
		var mult = n * i
		fmt.Println("the multiplication of", n, "and", i, "is", mult)
	}
}
