package main

import "fmt"

func main() {
	personas := make(map[string][]string)

	personas[`bond_james`] = []string{`shaken, no stirred`, `martinis`, `fast cars`}
	personas[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	personas[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	personas[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	
	for key, value := range personas {
		fmt.Printf("%s %s\n", key, value)
	}

	delete(personas, "no_dr")

	fmt.Println("Despues de borrar un registro: ----------------------------")
	for key, value := range personas {
		fmt.Printf("%s %s\n", key, value)
	}

}
