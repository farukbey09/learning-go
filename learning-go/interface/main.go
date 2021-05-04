package main

import (
	"fmt"
	"math"
)

type geometry interface {
	getPerimeter() float64
	getArea() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r *rect) getArea() float64 {
	return r.width * r.height
}

func (r *rect) getPerimeter() float64 {
	return 2 * (r.height + r.width)
}
func (c *circle) getArea() float64 {

	return math.Pi * c.radius * c.radius
}
func (c *circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g.getArea())
	fmt.Println(g.getPerimeter())
	fmt.Println("------")
}
func main() {
	rect1 := rect{
		width:  5,
		height: 3,
	}
	circle1 := circle{
		radius: 5,
	}

	measure(&rect1)
	measure(&circle1)

}
