package main

import (
	"fmt"
	"math"
)

func main() {

	var sec float64
	var min float64
	var hour float64

	ch := make(chan float64, 3)

	fmt.Println("Enter number of seconds: ")
	fmt.Scanln(&sec)

	ch <- sec

	min = sec / 60

	ch <- min

	hour = min / 60

	ch <- hour

	close(ch)

	fmt.Println("")

	for i := range ch {
		fmt.Println(math.Floor(i))
	}

}
