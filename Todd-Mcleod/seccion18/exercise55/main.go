package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make string
	model string
	doors int
	color string
}


func main() {

	vehicle1 := vehicle {
		engine: engine {
			electric: true,
		},
		make: "hierro",
		model: "R6",
		doors: 2,
		color: "red",
	}
	vehicle2 := vehicle {
		engine: engine {
			electric: false,
		},
		make: "plasstic",
		model: "R6343",
		doors: 4,
		color: "white",
	}

	fmt.Printf("%v\n", vehicle1)
	fmt.Printf("%v\n", vehicle2)
	fmt.Println("-------------------------------")
	fmt.Printf("el vehiculo 1 tiene un engine electrico? %t\n", vehicle1.engine.electric)
	fmt.Printf("el vehiculo 2 tiene un engine electrico? %t", vehicle2.engine.electric)
}