package main

import (
	"fmt"
)

type Person struct {
	name     string
	lastname string
	age      int
}

func (p *Person) speak() {
	fmt.Println("Hola mi nombre es ", p.name)
}

type Human interface {
	speak()
}

func saySomething(hum Human) {
	hum.speak()
}

func main() {
	fmt.Println("Esta es la funcion main")

	valentin := Person{
		name:     "Valentin",
		lastname: "Morardo",
		age:      22,
	}

	matias := Person{
		name:     "Matias",
		lastname: "Ramos",
		age:      23,
	}

	saySomething(&valentin)
	saySomething(&matias)
}
