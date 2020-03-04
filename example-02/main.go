package main

import "fmt"

type person struct {
	first string
}

type mongo map[int] person
type postgres map[int] person


// Both mongo and postgres can save and retrieve.
func (m mongo) save(n int, p person) {
	m[n] = p
}
func (m mongo) retrieve(n int) person {
	return m[n]
}
func (pg postgres) save(n int, p person) {
	pg[n] = p
}
func (pg postgres) retrieve(n int) person {
	return pg[n]
}

// accessor saying mongo and postgres
// being abstract types of itself.
type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

// an alternative way of implementing the service.
type personService struct {
	a accessor
}

// function for the get
func (ps personService) get(n int) (person, error) {
	p := ps.a.retrieve(n)
	if p.first == "" {
		return person{}, fmt.Errorf("no person with key %d\n", n)
	}
	return p, nil
}

// utilizing accessor as a param
// hence you can simply swap out the actual
// implementations at any point.
func put(a accessor, n int, p person) {
	a.save(n, p)
}
func get(a accessor, n int) person {
	return a.retrieve(n)
}


func main() {
	dbm := mongo{}
	dbp := postgres{}

	p1 := person {
		first: "James",
	}

	p2 := person {
		first: "Graham",
	}

	ps := personService{a:dbm}

	fmt.Printf("\nPrinting the dbm result\n\n")
	put(dbm, 1, p1)
	put(dbm, 2, p2)
	fmt.Println(get(dbm, 1))
	fmt.Println(get(dbm, 2))

	fmt.Println(ps.get(1))
	fmt.Println(ps.get(3))

	fmt.Printf("\n-------------\n\n")

	fmt.Printf("Printing the dbp result\n\n")

	put(dbp, 1, p1)
	put(dbp, 2, p2)
	fmt.Println(get(dbp, 1))
	fmt.Println(get(dbp, 2))
}
