package main

import "fmt"

func main() {
	/* for i := 1; i <= 4; i++ {
		for j := 4; j >= i; j-- {

			fmt.Print(" ")

		}
		for k := 1; k <= i; k++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}*/

	//Other Method

	for i := 1; i <= 4; i++ {
		for j := 4; j >= 1; j-- {

			if j > i {
				fmt.Print(" ")
			} else {
				fmt.Print(" *")
			}
		}
		fmt.Println("")
	}

}
