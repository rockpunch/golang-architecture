package main

import "fmt"

//this, we would say CONCRETE Type.
type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Printf("My name is %v\n", p.first)
}

func (sa secretAgent) speak() {
	fmt.Printf("I'm a secret agent %v\n", sa.first)
}

//an interface says
//this, we would say ABSTRACT Type.
//THUS, when you have a value of an interface type...
//you know NOTHING about WHAT IT IS, you know ONLY WHAT IT CAN DO.
type human interface {
	speak()
}

//foo takes abstract type of human, hence will take person, secretAgent type.
func foo (h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	sa1 := secretAgent{
		person: person{first:"Miss Moneypenny"},
		ltk:    true,
	}

	fmt.Printf("p1 is a type of %T\nand it can be represented as: %v\n:)\n\n", p1, p1)

	// in go a VALUE can be of more than one TYPE
	// in this example p1 is both (concrete)TYPE person and (abstract)TYPE human
	var h1, h2 human
	h1 = p1
	h2 = sa1

	// we know that human carries speak method
	h1.speak()

	// sa has person. logically, it has to have two speak method. but what would happen?
	h2.speak()

	fmt.Println("---------------")
	foo(h1)
	foo(h2)
	foo(p1)
	foo(sa1)
}