package main

import "fmt"

func main() {
	var hour = 0
	var min = 0
	var sec = 0

	for i := 1; i <= 120; i++ {
		sec++
		min++

		if sec == 60 || sec == 120 {
			min++
		}
		if min == 60 || min == 120 {
			hour++
		}

	}
	fmt.Println(sec)

	fmt.Printf("Min is : %d\n", min)
	fmt.Println("hour: ", hour)

	/*if min > 59 {
			hour++
		}

		if hour > 24 {
			fmt.Printf("Time is: %d %d,%d", hour, min, sec)
		}
	}*/

	//func (t Time) Clock() (hour, min, sec int)

}
