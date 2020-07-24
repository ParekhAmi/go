package main

import "fmt"

func main() {

	a := make(chan int, 3)

	a <- 1
	a <- 2
	a <- 3
	close(a)

	for i := range a {
		fmt.Println(i)
	}

}
