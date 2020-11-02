package main

import "fmt"

func main() {
	var a, b int = 4, 5
	var temp int

	temp = a
	a = b
	b = temp

	fmt.Println("Value of a: ", a)
	fmt.Println("Value of b: ", b)
}
