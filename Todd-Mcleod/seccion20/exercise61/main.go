package main

import "fmt"


type person struct {
	first string
	age int
}


func(p person) speak() {
	fmt.Printf("Hola. mi nombre es %s y tengo %d años", p.first, p.age)
} 

func main() {
	
	valentin := person {
		first: "Valentin",
		age: 22,
	}

	valentin.speak()
}


