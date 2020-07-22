package main

import "fmt"

func main() {
	//var hour = 0
	var min = 0
	var sec = 0

	for i := 1; i <= 60; i++ {
		sec++

		if sec > 59 {
			min++
		}

	}
	fmt.Println(sec)

	fmt.Printf("Min is : %d", min)

	/*if min > 59 {
			hour++
		}

		if hour > 24 {
			fmt.Printf("Time is: %d %d,%d", hour, min, sec)
		}
	}*/

}
