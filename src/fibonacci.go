package main

import "fmt"

func main() {

	var num int
	t1 := 0
	t2 := 1
	next := 0

	fmt.Print("Enter number of elements:")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
		}
		if i == 2 {
			fmt.Print(" ", t2)
		}
		next = t1 + t2
		t1 = t2
		t2 = next
		fmt.Print(" ", next)
	}

}
