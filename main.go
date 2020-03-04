package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Printf("My name is %v\n", p.first)
}

//an interface says

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
	}
	fmt.Printf("p1 is a type of %T\nand it can be represented as: %v\n:)\n", p1, p1)

	// in go a VALUE can be of more than one TYPE
	// in this example p1 is both (concrete)TYPE person and (abstract)TYPE human
	var h1 human
	h1 = p1
	fmt.Printf("\nh1 is a type of %T\n", h1)
}
