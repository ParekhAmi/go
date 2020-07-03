package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Declaration of Slice

	var slice1 []int                                  //Declare a slice
	var slice2 = []string{"India", "Canada", "Japan"} //Initialize using slice literal
	var slice3 = make([]string, 10, 10)               //Using make keyword
	var sliceNew = new([50]int)[0:10]                 //Using new keyword

	fmt.Println(slice1)
	fmt.Println(reflect.TypeOf(slice1))         // []int
	fmt.Println(reflect.ValueOf(slice1).Kind()) //slice
	fmt.Printf("Slice2 \t Len: %v \t Cap: %v\n", len(slice2), cap(slice2))
	fmt.Printf("Slice3 \t Len: %v \t Cap: %v\n", len(slice3), cap(slice3))

	fmt.Printf("sliceNew \tLen: %v \tCap: %v\n", len(sliceNew), cap(sliceNew))
	fmt.Println(sliceNew)
}
