package main

import (
	"fmt"
	architecture "github.com/rockpunch/golang-architecture"
	"github.com/rockpunch/golang-architecture/storage/mongo"
	"github.com/rockpunch/golang-architecture/storage/postgrs"
)

func main() {
	dbm := mongo.Db{}
	dbp := postgrs.Db{}

	p1 := architecture.Person{
		First: "James",
	}

	p2 := architecture.Person{
		First: "Graham",
	}

	ps := architecture.NewPersonService(dbm)

	fmt.Printf("\nPrinting the dbm result\n\n")
	architecture.PersonService.Put(dbm, 1, p1)
	architecture.PersonService.Put(dbm, 2, p2)
	fmt.Println(architecture.PersonService.Get(dbm, 1))
	fmt.Println(architecture.PersonService.Get(dbm, 2))

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	fmt.Printf("\n-------------\n\n")

	fmt.Printf("Printing the dbp result\n\n")

	architecture.PersonService.Put(dbp, 1, p1)
	architecture.PersonService.Put(dbp, 2, p2)
	fmt.Println(architecture.PersonService.Get(dbp, 1))
	fmt.Println(architecture.PersonService.Get(dbp, 2))
}

