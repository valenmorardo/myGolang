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
	fmt.Printf("%s %s tiene como gustos de helado fav los siguientes: \n", persona1.first_name, persona1.last_name)
	for _, v := range persona1.fav_iceCream {
		fmt.Printf("%s\n", v)
	}

	
	fmt.Printf("\n\n%s %s tiene como gustos de helado fav los siguientes: \n", persona2.first_name, persona2.last_name)
	for _, v := range persona2.fav_iceCream {
		fmt.Printf("%s\n", v)
	}

}