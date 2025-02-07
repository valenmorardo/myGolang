package main

import "fmt"

func main() {

	miSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slicing1 := miSlice[:5]
	slicing2 := miSlice[5:]
	slicing3 := miSlice[2:7]
	slicing4 := miSlice[1:6]

	fmt.Println(slicing1)
	fmt.Println(slicing2)
	fmt.Println(slicing3)
	fmt.Println(slicing4)

}