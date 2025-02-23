package main

import (
	"fmt"

)

type Persona struct {
	firstname string
	lastname  string
	age       int
}
func (p Persona) speak() {
	fmt.Printf("Yo soy %s", p.firstname)
}
func (p Persona) String() string{
	return fmt.Sprintf("Esta es la funcion Stringer de Persona, yo soy %s", p.firstname)
}



type Programador struct {
	Persona  Persona
	language string
}
func (p Programador) isProgramming(n1 int, n2 int) string {
	rta := n1 + n2
	fmt.Printf("Soy %s y estoy programando en %s.\n", p.Persona.firstname, p.language)
	return fmt.Sprintf("La respuesta de %d + %d es: %d", n1, n2, rta)
}
func (p Programador) speak() {
	fmt.Printf("Yo soy %s y programo en %s", p.Persona.firstname, p.language)
}
func (p Programador) String() string{
	return fmt.Sprintf("Esta es la funcion Stringer de Programador, yo soy %s", p.Persona.firstname)
}



type Humano interface {
	speak()
}


func saySomething(h Humano) {
	h.speak()
	fmt.Println("\nHolaaaa voy a decir soy humanoo")
}


func main() {
	valentin := Persona{
		firstname: "valentin",
		lastname:  "Morardo",
		age:       22,
	}

	matias := Persona{
		firstname: "Matias",
		lastname:  "ramos",
		age:       22,
	}

	valentinProgramador := Programador{
		Persona:  valentin,
		language: "Golang",
	}

	matiasProgramador := Programador{
		Persona:  matias,
		language: "Python",
	}

	fmt.Println(valentin, "---", matias)
	fmt.Println(valentinProgramador)

	fmt.Println(valentinProgramador.isProgramming(1, 2))
	fmt.Println(matiasProgramador.isProgramming(3, 1))
	
	valentinProgramador.speak()
	saySomething(valentin)
	saySomething(matiasProgramador)

	fmt.Println()

	fmt.Println(valentin)
	fmt.Println(matiasProgramador)




}
