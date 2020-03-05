package architecture

import (
	"testing"
)

type Db map[int]Person

func (m Db) Save(n int, p Person) {
	m[n] = p
}
func (m Db) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	mdb := Db{}
	p := Person {
		First: "James",
	}
	ps := PersonService{a:mdb}
	ps.Put(1, p)
	result, err := ps.Get(1)
	if err != nil {
		t.Fatalf("Failed due to: %v", err)
	}
	if result != p {
		t.Fatalf("Want %v, got %v", p, result)
	}
}