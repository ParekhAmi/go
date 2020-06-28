package main

import "fmt"

var num int
var fact int = 1

func factorial(a int) int {
	if num < 0 {

		fmt.Println("It's negative number, please enter positive number.")

	} else {

		for i := a; i > 0; i-- {

			fact *= i
		}
	}
	return fact
}

func main() {

	fmt.Println("Enter Positive Number: ")
	fmt.Scanln(&num)

	factNum := factorial(num)
	fmt.Println(factNum)

}
