package main

import (

	"fmt"
	"math/rand"

)

func main() {


	for x := 0; x < 5; {
		if(rand.Intn(2) == 1) {
			x++
		} else {
			x--
		}
		fmt.Println(x)
	}
}