package main

import "fmt"

func main() {
	personas := make(map[string][]string)

	personas[`bond_james`] = []string{`shaken, no stirred`, `martinis`, `fast cars`}
	personas[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	personas[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}


	
	for key, value := range personas {
		fmt.Printf("LA persona %s tiene estos gustos\n", key)
		for i, value2 := range value {
			fmt.Printf("%s ---> posicion %d \n", value2, i)
		}
		fmt.Println("-------------------------------------------")
	}

}
