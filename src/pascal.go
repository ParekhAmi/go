package main

import (
	"fmt"
)

func main() {

	var row, col, val int
	var number int = 6

	for row = 1; row <= number; row++ {
		for col = 1; col <= row; col++ {

			if row == col || col == 1 {
				fmt.Print("1")
			} else {

				val = col * (row - col + 1) / col

				fmt.Print(val)
				//fmt.Print("*")
			}

		}

		fmt.Println("")
	}

}

/*
	var row, col float64
	var maxColLimit, maxRowLimit float64 = 7, 4
	var halfLen float64 = math.Ceil(maxColLimit / 2)

	for row = 1; row <= maxRowLimit; row++ {
		i := 0

		for col = 1; col <= maxColLimit; col++ {

			if col < (halfLen - row + 1) {
				fmt.Print(" ")
			} else {
				if col > (halfLen + row - 1) {
					fmt.Print(" ")
				} else {
					i++

					if i%2 == 0 {
						fmt.Print(" ")
					} else {
						if col == (halfLen-row+1) || col == (halfLen+row-1) {

							fmt.Print("1")

						} else {
							<= halfLen {

							}
							fmt.Print("3")
						}

					}
				}
			}

		}

		fmt.Println("")
	}
}
*/
