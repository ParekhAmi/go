package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Append items using append function

	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20

	fmt.Println("Slice A: ", a)

	a = append(a, 30, 40, 50, 60, 70, 80)
	fmt.Println("Slice B: ", a)
	fmt.Printf("Slice B  Len:%v \t Cap:%v\n", len(a), cap(a))

	fmt.Println("--------------------------")

	//Access items

	intSlice := []int{10, 20, 30, 40}

	fmt.Println(intSlice[0])
	fmt.Println(intSlice[1])
	fmt.Println(intSlice[0:4])

	fmt.Println("--------------------------")

	//Change item value

	intSlice[3] = 50

	fmt.Println(intSlice)

	fmt.Println("--------------------------")

	//Delete an element

	var countryValue = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(countryValue)

	countryValue = removeIndex(countryValue, 3)
	fmt.Println(countryValue)

	fmt.Println("--------------------------")

	//copy one slice items into another slice

	c := []int{5, 6, 7}
	d := make([]int, 5, 10) // Create a bigger slice

	copy(d, c)
	fmt.Println("Slice B after copying:", d)

	d[3] = 8
	d[4] = 9
	fmt.Println("Slice B after adding elements:", d)

	fmt.Println("--------------------------")

	//Iterate over a slice

	var country = []string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Println("***************")
	for index, element := range country {
		fmt.Println(index, "--", element)
	}

	fmt.Println("***************")
	for _, value := range country {
		fmt.Println(value)
	}

	fmt.Println("***************")
	m := 0
	for range country {
		fmt.Println(country[m])
		m++
	}

	fmt.Println("--------------------------")

	//Append a slice to an existing slice

	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}

	slice3 := append(slice2, slice1...)
	fmt.Println(slice3)

	fmt.Println("--------------------------")

	//Check itemExists in slice

	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strSlice, "Canada"))
	fmt.Println(itemExists(strSlice, "Africa"))

}

func removeIndex(slice []string, index int) []string {
	return append(slice[0:index], slice[index+1:]...)
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
