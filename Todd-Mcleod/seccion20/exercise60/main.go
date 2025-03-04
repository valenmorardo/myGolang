package main

import "fmt"


func main() {
	
	defer fmt.Println("0 print")
	fmt.Println("primer print")

	defer fmt.Println("Segundo print")

	fmt.Println("tercer print")
	fmt.Println("cuarto print")
	fmt.Println("quinto print")
	defer fmt.Println("sexto  print")
	defer fmt.Println("septimo  print")
}


