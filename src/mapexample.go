package main

import (
	"fmt"
	"sort"
)

func main() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20} //Initialization
	// var employee = map[string]int{} //empty map declaration

	fmt.Println(employee)

	fmt.Println("---------------")

	var employeeMake = make(map[string]int) //Declaration using make function
	employeeMake["Ami"] = 10
	employeeMake["Mitul"] = 20

	fmt.Println(employeeMake)

	fmt.Println("---------------")

	//Accessing item
	fmt.Println(employeeMake["Mitul"])

	fmt.Println("---------------")

	//Add item
	employeeMake["@mi"] = 30
	fmt.Println(employeeMake)

	//Delete Function deletes item from map
	delete(employeeMake, "Ami")
	fmt.Println(employeeMake)

	fmt.Println("---------------")

	//Iterate over a map
	var members = map[string]int{"Mark": 10, "Sandy": 20, "Rocky": 30, "Rajiv": 40, "Kate": 50}

	for key, element := range members {
		fmt.Println("Key: ", key, " => ", "Element: ", element)
	}
	fmt.Println("---------------")

	//Truncate map - All items will be cleared... Two methods
	var student = map[string]int{"Mark": 10, "Sandy": 20, "Rocky": 30, "Rajiv": 40, "Kate": 50}

	fmt.Println(student)
	// Method - I
	for k := range student {
		fmt.Println(student[k])
		delete(student, k)
		fmt.Println(student[k])

	}

	// Method - II
	//student = make(map[string]int)
	fmt.Println("---------------")

	//Sort map keys

	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	keys := make([]string, 0, len(unSortedMap))

	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}

	fmt.Println("---------------")

	//Sort map values

	values := make([]int, 0, len(unSortedMap)) // Int slice to store values of map.

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	sort.Ints(values) // Sort slice values.

	for _, v := range values {
		fmt.Println(v) // Print values of sorted Slice.
	}
	fmt.Println("---------------")

	//Merge maps

	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first)

}
