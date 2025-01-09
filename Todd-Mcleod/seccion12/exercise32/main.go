package main

import (
	"fmt"
)

func main() {
	var x int

	for {
		fmt.Println(x)
		if x == 10 {
			break
		}
		x++
	}
}
