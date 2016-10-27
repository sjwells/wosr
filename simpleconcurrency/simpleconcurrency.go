package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello\n")
	go slowFunction()
	fmt.Printf("Bye!\n")
}

func slowFunction() {
	//do something slow
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished slow stuff\n")
}
