package main

import (
	"fmt"

)

type customErr struct {
	extraInfo string

}

func (ce customErr) Error() string {
	return fmt.Sprintf("Ocurrio un error. ERROR: %v", ce.extraInfo)
	
}

func main() {

	c1 := customErr{
		extraInfo: "el error es este X",
	}

	foo(c1)

}

func foo(e error) {
	fmt.Println(e)
}
