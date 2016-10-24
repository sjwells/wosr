package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	x, y := split(15)
	fmt.Printf("x=%d, y=%d\n", x, y) //outputs x=6, y=9
}
