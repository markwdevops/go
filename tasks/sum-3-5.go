package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			var sum = n + i
			fmt.Println("The addition of", i, "and", n, "is", sum)
		}

	}
}
