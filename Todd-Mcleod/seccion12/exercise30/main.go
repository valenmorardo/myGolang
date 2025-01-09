package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {

		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("Vuelta %v:  X es 0 \n", i)
		case 1:
			fmt.Printf("Vuelta %v:  X es 1 \n", i)
		case 2:
			fmt.Printf("Vuelta %v:  X es 2 \n", i)
		case 3:
			fmt.Printf("Vuelta %v:  X es 3 \n", i)
		case 4:
			fmt.Printf("Vuelta %v:  X es 4 \n", i)
		default:
			fmt.Printf("Vuelta %v:  x no esta entre 0 y 5 \n", i)
		}

	}
}
