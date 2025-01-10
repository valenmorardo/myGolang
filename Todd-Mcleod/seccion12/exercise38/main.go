package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i <= 100; i++ {
		if x := rand.Intn(6); x == 3 {
			fmt.Printf("X vale %v \n", x)
		}
	}
}
