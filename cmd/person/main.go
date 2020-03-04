package main

import (
	"fmt"
	"github.com/rockpunch/golang-architecture"
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
	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbm, 2))

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	fmt.Printf("\n-------------\n\n")

	fmt.Printf("Printing the dbp result\n\n")

	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))
}

