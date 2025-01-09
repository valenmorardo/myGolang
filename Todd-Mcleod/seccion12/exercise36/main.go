package main

import (
	"fmt"
)

func main() {

	datos := map[string]int {
		"James": 42,
		"Moneypenny": 32,
	}

	for key, value := range datos {
		fmt.Printf("%s tiene %v años \n", key, value)
	}

}