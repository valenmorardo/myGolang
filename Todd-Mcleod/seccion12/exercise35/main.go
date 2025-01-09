package main

import (
	"fmt"
)

func main() {

	x := []int{42, 43, 44, 45, 46, 47}
	
	for index, value := range x {
		
		fmt.Printf("En la posicion %v esta el numero %v \n", index, value)
	}
}