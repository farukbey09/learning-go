package main

import "fmt"

type edges struct {
	x, y int
}

func (e *edges) GetPerimeter() int {
	return 2 * (e.x + e.y)
}

func (e *edges) GetArea() int {
	return e.x * e.y
}

func main() {
	edge := edges{2, 5}
	fmt.Println(edge.x)
	fmt.Println(edge.y)
	fmt.Println(edge.GetPerimeter())
	fmt.Println(edge.GetArea())

	fmt.Println("-------------")

	edge2 := &edge
	fmt.Println(edge2.x)
	fmt.Println(edge2.y)
	fmt.Println(edge2.GetPerimeter())
	fmt.Println(edge2.GetArea())

}
