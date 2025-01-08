package main

import (
	"fmt"
	"math/rand"
)

func randomInt(n int) int {
	return rand.Intn(n)
}

func main() {
	x := randomInt(3)
	fmt.Printf("El valor de la variable x es: %v y es de tipo %T \n", x, x)

	switch {
	case x <= 100:
		fmt.Printf("El numero %v de la variable x esta entre 0 y 100", x)

	case x >= 101 && x <= 200:
		fmt.Printf("El numero %v de la variable x esta entre 101 y 200", x)
	case x >= 201 && x <= 250:
		fmt.Printf("El numero %v de la variable x esta entre 201 y 250", x)

	}
}
