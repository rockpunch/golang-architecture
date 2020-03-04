package architecture

import (
	"fmt"
)

type Person struct {
	First string
}

// Accessor saying mongo and postgrs
// being abstract types of itself.
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

// an alternative way of implementing the service.
type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{a:a}
}

// function for the get
func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("no Person with name %d\n", n)
	}
	return p, nil
}

func (ps PersonService) Put(n int, p Person) (Person, error) {
	if p.First == "" {
		return Person{}, fmt.Errorf("person must contain First value\n")
	}
	return ps.Put(n, p)
}