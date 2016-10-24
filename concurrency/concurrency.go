package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	responses := make(chan string)
	for i := 1; i <= 5; i++ {
		go slowFunction(i, responses)
	}
	fmt.Printf("Winner was %s\n", <-responses)
}

func slowFunction(num int, out chan<- string) {
	waitSeconds := rand.Intn(10)
	fmt.Printf("Num %d is waiting %d seconds\n", num, waitSeconds)
	time.Sleep(time.Duration(waitSeconds) * time.Second)
	out <- fmt.Sprintf("%d", num)
}
