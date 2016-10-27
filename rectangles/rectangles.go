package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func Area(r Rectangle) int {
	return r.width * r.height
}

func main() {
	r := Rectangle{width: 10, height: 5}
	fmt.Printf("area: %d\n", r.Area())
	fmt.Printf("area as a function: %d\n", Area(r))
}
