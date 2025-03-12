package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Println("VALUES (memory adress)")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	fmt.Println()
	fmt.Println("TYPES")
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", c)
	fmt.Printf("d: %T\n", d)
	fmt.Println()
	fmt.Println("Data storage at memory locations")
	fmt.Println("a: ", *a)
	fmt.Println("b: ", *b)
	fmt.Println("c: ", *c)
	fmt.Println("d: ", *d)
	
}
