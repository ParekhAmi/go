package main

import "fmt"

func main() {
	var i int

	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
	}()
	i = <-ch
	j := <-ch
	fmt.Println(i)
	fmt.Println(j)
}
