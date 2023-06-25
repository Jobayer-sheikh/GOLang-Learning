package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	x, y float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.x * r.y
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.x + r.x)
}

func main() {
	var val [3]interface{}
	val[0] = 1
	val[1] = "string"
	val[2] = 22 / 7
	fmt.Println(val)

	r := rectangle{x: 10, y: 20}
	c := circle{radius: 30}

	var rec geometry = r
	fmt.Println(rec.area())
	fmt.Println(rec.perimeter())

	var cir geometry = c
	fmt.Println(cir.area())
	fmt.Println(cir.perimeter())
}
