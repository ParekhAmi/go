package main

import "fmt"

var a, b int
var add, sub, mul int
var div, mod float64

func arithmetic() {

	add = a + b
	sub = a - b
	mul = a * b
	div = float64(a / b)
	mod = float64(a % b)

}

func main() {
	fmt.Println("Enter two numbers: ")
	fmt.Scanln(&a, &b)

	arithmetic()

	fmt.Println("Addition is: ", add)
	fmt.Printf("Subtraction is: %d\n", sub)
	fmt.Printf("Multiplication is: %d\n", mul)
	fmt.Println("Division is: ", div)
	fmt.Println("Modulo is :", mod)
}
