package main

import "fmt"

// var a, b = 2, 2

func arithmetic() {

	var a, b int
	var sum, sub, mul, mod int

	fmt.Print("Enter any two numbers: ")
	fmt.Scan(&a, &b)

	sum = a + b
	sub = a - b
	mul = a * b
	mod = a % b

	fmt.Print("Sum : ", sum)
	fmt.Print("\nSub : ", sub)
	fmt.Print("\nMul : ", mul)
	fmt.Print("\nMod : ", mod)
}

func main() {

	arithmetic()

}
