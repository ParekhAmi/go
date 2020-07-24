package main

import (
	"fmt"
	"math"
)

func main() {

	/* fmt.Println("-----------First method----------")

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
	}*/

	// fmt.Println("-----------second method----------")

	var sec, second float64
	var min, minute float64
	var hour, hours float64

	ch := make(chan float64, 3)

	fmt.Println("Enter number of seconds: ")
	fmt.Scanln(&sec)

	ch <- sec

	second = <-ch

	min = sec / 60

	ch <- min

	minute = <-ch

	hour = min / 60

	ch <- hour

	hours = <-ch

	fmt.Println(math.Floor(hours), ":", math.Floor(minute), ":", math.Floor(second))

}
