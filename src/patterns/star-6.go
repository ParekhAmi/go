package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 8; j++ {
			if j <= 5-i || j >= 4+i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
