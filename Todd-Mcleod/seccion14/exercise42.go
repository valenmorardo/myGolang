package main

import "fmt"

func main() {

	miArray := []int{1, 2, 3, 4, 5}

	for index, value := range miArray {
		fmt.Printf("El numero de la posicion: %v es %v y es de tipo %T \n", index, value, value)
	}
}