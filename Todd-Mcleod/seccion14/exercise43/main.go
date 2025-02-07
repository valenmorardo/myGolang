package main

import "fmt"

func main() {

	miSlice := []int{42, 43, 44, 45, 45, 47, 48, 49, 50, 51}

	for _, value := range miSlice {
		fmt.Printf("%v - %T\n", value ,value)
	}
}