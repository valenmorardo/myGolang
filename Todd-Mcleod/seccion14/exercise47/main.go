package main

import "fmt"

func main() {
	provincias := make([]string, 0, 5) 

	fmt.Printf("capacidad: %v\n", cap(provincias))
	provincias = append(provincias, "Cordoba", "mendoza", "San luis", "Santa Fe", "Neuquen", "Entre Rios")
	fmt.Println(provincias)
	fmt.Printf("capacidad: %v\n", cap(provincias))
	fmt.Printf("Length: %v\n", len(provincias))

	for i := 0; i < len(provincias); i++ {
		fmt.Printf("En la posicion %v esta la provincia %s\n", i, provincias[i])
	}


}
