package main

import "fmt"

func main() {
	var sum, n int

	fmt.Println("Enter natural number: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
