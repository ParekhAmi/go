package main

import "fmt"

func main() {
	var a int = 2
	var p *int = &a

	fmt.Println(a)
	fmt.Println(p)

	*p = 54

	fmt.Println(a)

	var pp = &p

	fmt.Println(pp)
	fmt.Println(*pp)
	fmt.Println(**pp)

}
