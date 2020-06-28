package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	b := [5]string{"Ami", "Mitul", "Krushna", "Hari", "Hima"}
	c := [...]float64{1.2, 2.3, 3.4, 4.5, 5.5, 6.5}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//two dimensional

	var d [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i + j
		}
	}
	fmt.Println(d)
}
