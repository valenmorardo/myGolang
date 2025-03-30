package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Esta es la funcion main")

	wg.Add(2)
	go sayHello()
	go sayBye()

	wg.Wait()

}

func sayHello() {
	fmt.Println("Holaaaa")
	wg.Done()
}

func sayBye() {
	fmt.Println("Chauu")
	wg.Done()
}
