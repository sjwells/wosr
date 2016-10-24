package main

import "fmt"

type Rect struct {
	width  int
	height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func area(r Rect) int {
	return r.width * r.height
}

func main() {
	r := Rect{width: 10, height: 5}
	fmt.Printf("area: %d\n", r.area())
	fmt.Printf("area as a function: %d\n", area(r))
}
