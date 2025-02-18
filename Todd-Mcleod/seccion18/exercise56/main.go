package main

import "fmt"



func main() {

	persona := struct {
		first string
		friends map[string]int
		favDrinks []string
	} {
		first: "valentin",
		friends: map[string]int{
			"Tutin": 25,
			"Mati": 23,
			"niufa": 23,
		},
		favDrinks: []string{"Vodka", "Vino", "Coca"},
	}

	fmt.Printf("Los amigos y tragos favoritos de %s son:\n", persona.first)
	fmt.Println("AMIGOS:")
	for amigo, edad := range persona.friends {
		fmt.Printf("%s tiene %d\n", amigo, edad)
	}
	fmt.Println("TRAGOS:")
	for _, drink := range persona.favDrinks {
		fmt.Printf("%s\n", drink)
	}




}