package main

import "fmt"

func main() {
	fmt.Println("Enter your name")
	var name string
	fmt.Scanln(&name)

	if name == "Alice" || name == "Bob" {
		fmt.Println("Hi", name)
	}
}
