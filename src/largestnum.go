package main

import "fmt"

func main() {
	var arr [100]int
	var temp int

	fmt.Println("Enter number of elements: ")
	fmt.Scanln(&temp)

	for i := 0; i < temp; i++ {
		fmt.Println("Enter number into an array")
		fmt.Scanln(&arr[i])

	}
	for j := 1; j < temp; j++ {
		if arr[0] < arr[j] {
			arr[0] = arr[j]
		}
	}
	fmt.Print("largest number is: ", arr[0])
}
