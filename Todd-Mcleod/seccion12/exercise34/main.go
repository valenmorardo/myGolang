package main

import (
	"fmt"
)

func main() {
	
	for x := 0; x <=5; x++ {
		
		for i:=0; i <= 5; i++ {
			fmt.Printf("EXTERIOR Loop: Vuelta %v \t INTERIOR Loop. Vuelta %v \n", x, i)
		}
	} 
}