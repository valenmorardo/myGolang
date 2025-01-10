package main

import (
	"fmt"
)

func main() {
	datos := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	if age, ok := datos["James"]; ok {
		fmt.Printf("La edad es %v", age)
	} else {
		fmt.Println("No existe el valor buscado")
	}
}
