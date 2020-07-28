package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length, width float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return math.Ceil(r.length * r.width)
}

func (r rectangle) perimeter() float64 {
	return math.Ceil(2 * (r.length + r.width))
}
func (c circle) area() float64 {
	return math.Ceil(math.Pi * c.radius * c.radius)
}
func (c circle) perimeter() float64 {
	return math.Ceil(2 * math.Pi * c.radius)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {

	r := rectangle{2, 4}
	c := circle{3}

	measure(r)
	measure(c)
}
