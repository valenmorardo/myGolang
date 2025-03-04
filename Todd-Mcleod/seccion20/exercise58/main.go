package main

import "fmt"


func main() {
	x := foo()

	i, s := bar()

	fmt.Println(x)
	fmt.Println(i, s)

}



func foo() int {

	return 2
}

func bar() (int, string) {
	return 42, "Hola"
}