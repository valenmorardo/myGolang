
package main

import "fmt"


func main() {

	sliceHelados := []string{"Chocolate", "Vainilla", "Menta granizada", "Granizado", "Frutilla"}

	fmt.Println(sliceHelados)
	fmt.Printf("El tipo de slice helados es %T", sliceHelados)


	for _, value := range sliceHelados {
		fmt.Println(value)
	}
}