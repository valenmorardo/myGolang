package main

import (
	"fmt"
)

func main() {
	for x := 0; x <= 100; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Printf("%v es impar \n", x)

	}
}
