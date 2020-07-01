package main

import "fmt"

func main() {

	var arrayTwo [3][4]int
	var i, j int

	for i = 0; i < 3; i++ {
		for j = 0; j < 4; j++ {
			fmt.Print("Enter value for two dimensional array: ")
			fmt.Scanf("%d\n", &arrayTwo[i][j])
		}

	}
	for i = 0; i < 3; i++ {
		for j = 0; j < 4; j++ {

			fmt.Printf(" %d ", arrayTwo[i][j])

			if j == 3 {
				fmt.Println("")
			}
		}

	}

}
