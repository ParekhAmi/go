package main

import "fmt"

func main() {

	var arrayTwo [3][4]int
	var i, j int

	for i = 0; i < 3; i++ {
		for j = 0; j < 4; j++ {
			fmt.Print("Enter value for two dimensional array: ")
			fmt.Scan(&arrayTwo[i][j])

		}
		fmt.Println(arrayTwo)
		fmt.Println("")
	}

}
