package main

import (
	"fmt"
)

func main() {

	var sec int
	var min *int
	var hour **int

	min = &sec
	hour = &min

	fmt.Println("Enter number of seconds: ")
	fmt.Scanln(&sec)

	*min = sec / 60

	**hour = *min / 60

	fmt.Println(**hour)
	fmt.Println(*min)
	fmt.Println(sec)

	/* var hour = 0
	var min = 0
	var sec = 0

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

	fmt.Printf("Min is : %d\n", min)
	fmt.Println("hour: ", hour)

	if hour > 24 {

		fmt.Printf("Time is: %d %d,%d", hour, min, sec)
	}*/

}
