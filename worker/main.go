package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("Worker running...")
		time.Sleep(2 * time.Second)

		// simulate random crash
		if rand.Intn(10) > 7 {
			fmt.Println("Simulating failure...")
			panic("Random crash")
		}
	}
}
