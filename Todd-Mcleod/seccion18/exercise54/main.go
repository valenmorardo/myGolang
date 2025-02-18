package main

import "fmt"

type person struct {
	first_name string
	last_name string
	fav_iceCream []string
}

func main() {

	persona1 := person{
		first_name: "valentin",
		last_name: "Morardo",
		fav_iceCream: []string{"Menta granizada", "Granizado", "Limon"},
	}

	persona2 := person{
		first_name: "Tutin",
		last_name: "Ontivero",
		fav_iceCream: []string{"dulce de leche", "chocolate", "frutilla"},
	}

	mapaPersonas := make(map[string]person)
	
	mapaPersonas[persona1.last_name] = persona1
	mapaPersonas[persona2.last_name] = persona2

	
	for k, v := range mapaPersonas {
		fmt.Printf("\nLos helados favoritos de %s son\n", k)
		for _, v2 := range v.fav_iceCream {
			fmt.Printf("%s \n", v2)
		}
	}


}