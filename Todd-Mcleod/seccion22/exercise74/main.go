package main

import "fmt"

func main() {
	a := 24

	y := &a

	fmt.Println(*y)
	fmt.Println(a)
}
