package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Germany", "Ger"))
	fmt.Println(strings.Contains("Germany", "er"))
	fmt.Println(strings.Contains("Germany", "sy"))

	fmt.Println(strings.ContainsAny("Germany", "G"))
	fmt.Println(strings.ContainsAny("Germany", "g"))

	fmt.Println(strings.Count("Cheese", "e"))

	fmt.Println(strings.EqualFold("Germany", "GERMANY"))
	fmt.Println(strings.EqualFold("Germany", "Germanyan"))

}
