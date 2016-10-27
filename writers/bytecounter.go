package main

import (
	"fmt"
)

type ByteCounter struct {
	counter int
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	(*c).counter += len(p)
	return len(p), nil
}

func main() {
	bc := ByteCounter{}
	fmt.Fprintf(&bc, "count these bytes")
	fmt.Printf("ByteCounter counted %d bytes\n", bc.counter)
}
