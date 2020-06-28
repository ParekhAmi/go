package main

import (
	"fmt"
	"math"
)

func main() {
	var rad, area, circum float64

	fmt.Print("Enter radius: ")
	fmt.Scan(&rad)

	area = math.Pi * math.Pow(rad, 2)
	circum = 2 * math.Pi * rad

	fmt.Println("Area of circle is: ", area)
	fmt.Println("circumference of circle is: ", circum)

}
