package main

import (
	"fmt"
	"sort"
)

func main() {

	// Ints() Function for Integer Slice sorting

	intSlice := []int{10, 5, 25, 8, 35, 20}

	sort.Ints(intSlice)
	fmt.Println("Integer slice after sort in ascending order: ", intSlice)

	fmt.Println("--------------------")

	//Strings() Function for String Slice sorting

	strSlice := []string{"Japan", "America", "France", "India"}

	sort.Strings(strSlice)
	fmt.Println("String slice after sort in ascending order: ", strSlice)

	fmt.Println("--------------------")

	// Capital String Slice
	strSlice = []string{"JAMAICA", "Estonia", "indonesia", "hong Kong"} // unsorted
	fmt.Println("Slice of string BEFORE sort:", strSlice)
	sort.Strings(strSlice)
	fmt.Println("Slice of string AFTER  sort:", strSlice)
}
