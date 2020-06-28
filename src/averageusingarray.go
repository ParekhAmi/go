package main

import "fmt"

func main() {
	var arr [5]int
	var sum int
	var avg float64

	for i := 0; i < len(arr); i++ {
		fmt.Print("Enter element: ")
		fmt.Scanln(&arr[i])

		sum += arr[i]
	}
	fmt.Println(arr)

	avg = float64(sum / len(arr))

	fmt.Println("Average of an array: ", avg)

}
