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
		return Person{}, fmt.Errorf("no Person with key %d\n", n)
	}
	return p, nil
}

// utilizing Accessor as a param
// hence you can simply swap out the actual
// implementations at any point.
func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}
func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}
