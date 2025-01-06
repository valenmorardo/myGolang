package main

import "fmt"

var name string = "Valentin"

const surname string = "Morardo"

func main() {
	age := 22

	fmt.Printf("Mi nombre es %s %s y mi edad es %v", name, surname, age)
}