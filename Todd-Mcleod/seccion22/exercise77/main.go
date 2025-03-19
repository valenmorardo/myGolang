package main


import "fmt"



type Person struct {
	first string
}




func main() {
	valentin := Person{
		first: "Valentin",
	}
	fmt.Println(valentin.first)
	valentin = changeNameSem(valentin, "mati")
	fmt.Println(valentin.first)

	changeNamePointer(&valentin, "Julian")
	fmt.Println(valentin.first)

}


func changeNameSem(p Person, s string) Person {
	p.first = s
	return p
}

func changeNamePointer(p *Person, s string) {
	p.first = s
}