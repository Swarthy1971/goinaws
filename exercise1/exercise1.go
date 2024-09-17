package main

import "fmt"

const PI = 3.14

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		radius: radius,
	}
}

func (t Circle) circumference() {
	circumference := 2 * PI * t.radius
	fmt.Println(circumference)
}

func (t Circle) Area() {
	fmt.Println(PI * t.radius * t.radius)
}

func main() {

	myCircle := NewCircle(1.5)
	myCircle.circumference()
	myCircle.Area()
}
