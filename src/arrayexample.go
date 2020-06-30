package main

import "fmt"

func main() {

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by refrence
	//Array declaration with non-fixed length
	var Array = [...]int{1, 2, 3, 4, 5, 6}

	//Static Array declaration with specific array index
	var arrayIndex = [5]int{2: 25, 4: 50}

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)

	//strArray1[0] = "Canada"
	strArray3[0] = "Canada"

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("*strArray3: %v\n", *strArray3)

	fmt.Println(len(Array))
	fmt.Println(arrayIndex)
}
