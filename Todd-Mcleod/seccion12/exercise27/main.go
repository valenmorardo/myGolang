package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	y := rand.Intn(11)

	fmt.Printf("El valor de X es %v \n", x)
	fmt.Printf("El valor de Y es %v \n", y)

	if (x < 4) && (y < 4) {
		fmt.Println("X e Y son menores que 4.")
	} else if (x > 6) && (y > 6) {
		fmt.Println("X e Y son mayores que 6")
	} else if (x >= 4) && (x <= 6) {
		fmt.Println("X es mayor o igual a 4 pero menor o igual a 6")
	} else if y != 5 {
		fmt.Println("Y no es 5")
	} else {
		fmt.Println("ningun caso de los anteriores es el actual")
	}
}
