package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println("")

	}
	for i := 1; i <= 4; i++ {
		for k := 7; k >= i; k-- {
			fmt.Print(" ")
		}
		for l := 1; l <= i; l++ {
			fmt.Print("*")
		}
		fmt.Println("")

	}

}
