package main

import "fmt"

func main() {
	var first, second, full string

	fmt.Println("Enter first string: ")
	fmt.Scan(&first)

	fmt.Println("Enter second string: ")
	fmt.Scan(&second)

	full = first + second

	fmt.Print(full)
}
