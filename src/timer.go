package main

import (
	"fmt"
)

func main() {

	var sec float64
	var min *float64
	var hour **float64

	min = &sec
	hour = &min

	fmt.Println("Enter number of seconds: ")
	fmt.Scanln(&sec)

	fmt.Println("Seconds are: ", sec)

	*min = sec / 60

	fmt.Println("Minutes are: ", *min)

	**hour = *min / 60

	fmt.Println("Hours are: ", **hour)

	fmt.Println(**hour, ":", *min, ":", sec)

}
