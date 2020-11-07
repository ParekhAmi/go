package main

import (
	"fmt"
	"math"
)

func main() {
	var str string
	var mid float64
	var midValue int
	palindrome := false

	fmt.Printf("Enter the value of string: ")
	fmt.Scanln(&str)

	fmt.Println("Length of string: ", len(str))

	mid = float64(len(str)) / 2
	midValue = int(math.Ceil(mid))
	fmt.Println("Midvalue : ", midValue)

	for l := 0; l < midValue; l++ {
		if str[l] != str[len(str)-(l+1)] || len(str) <= 2 {
			palindrome = false
			break
		} else {
			palindrome = true
		}

	}

	if palindrome == true {
		fmt.Println(str, " String is  a palindrome")
	} else {
		fmt.Println(str, " String is not a palindrome")
	}

	//Getting characters from string
	/* // #1 Method
	ch := strings.Split(str, "")
	fmt.Println(ch)
	for j := 0; j < len(str); j++ {
		fmt.Print(ch[j], " ")
	}

	fmt.Println("")

	// #2 Method
	for i, char := range str {
		fmt.Printf("%d  %c\n", i, char)

	}*/

}
