package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 2; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := 4; k >= i; k-- {
			fmt.Print("*")
		}

		fmt.Println("")
	}
}
