package main

import (
	"fmt"
	"math"
)

func main() {

	var row, col float64
	var maxColLimit, maxRowLimit float64 = 15, 8
	var halfLen float64 = math.Ceil(maxColLimit / 2)

	for row = 1; row <= maxRowLimit; row++ {
		i := 0
		for col = 1; col <= maxColLimit; col++ {

			if col < (halfLen + 1 - row) {

				fmt.Print(" ")

			} else {

				if col > (halfLen + row - 1) {

					fmt.Print(" ")

				} else {

					i++

					if i%2 == 0 {
						fmt.Print(" ")
					} else {
						fmt.Print("*")
					}

				}

			}

		}
		fmt.Println("")
	}
}

/*
func main() {
	var rows, cols int

	fmt.Println("Enter number of rows: ")
	fmt.Scanln(&rows)

	for i := 1; i <= rows; i++ {

		for cols = 1; cols <= 2*rows-1; cols++ {

			if cols > rows-(i-1) && cols < rows+(i+1) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}

		fmt.Println("")
	}

}
*/
