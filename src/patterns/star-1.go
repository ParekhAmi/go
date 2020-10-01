package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Print("*")
			} else {
				fmt.Print("*")
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
