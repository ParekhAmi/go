package main

import (
	"fmt"
	"math"
)

func main() {

	var hour float64 = 0
	var min float64 = 0
	var sec float64 = 0

	for i := 1; i <= 60; i++ {
		sec++

		if sec == 60 {
			min++

			sec = 0
		}

	}
	for j := 1; j <= 60; j++ {
		min++

		if min == 60 {
			hour++
			min = 0
		}
	}
	fmt.Println(sec)

	fmt.Printf("Min is : %.2f\n", min)
	fmt.Println("hour: ", hour)

	fmt.Println(math.Floor(hour), math.Floor(min), math.Floor(sec))

}
