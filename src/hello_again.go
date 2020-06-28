package main

import (
	"fmt"
	"strconv"
)

var a, b float64
var c string

func op(a, b float64) string {
	switch c {
	case "+":
		return strconv.FormatFloat(a+b, 0, 64, 64)
	case "-":
		return strconv.FormatFloat(a-b, 0, 64, 64)
	default:
		return "Invalid"
	}
}

func main() {

	fmt.Println("Enter two values: ")
	fmt.Scanln(&a, &b)

	fmt.Println("Enter operation: ")
	fmt.Scanln(&c)

	s := op(a, b)
	fmt.Println(s)
}
