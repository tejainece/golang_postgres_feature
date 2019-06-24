package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Address struct {
	Street string
	City string
}

func (ad *Address) Scan(value interface{}) error {
	fmt.Println(string(value.([]byte)))
	_, err := fmt.Sscanf(string(value.([]byte)), "(%s,%s)", &ad.Street, &ad.City)
	return err
}

type Person struct {
	ID string
	Name string
	Address Address
}

func main()  {
	var err error
	var db *sql.DB
	if db, err = sql.Open("postgres", "postgres://learner:thought@localhost:5432/learn?sslmode=disable"); err != nil {
		panic(errors.Wrapf(err, "error opening database"))
	}
	
	var postgresVersion string
	if err = db.QueryRow("SELECT version()").Scan(&postgresVersion); err != nil {
		panic(err)
	}
	
	fmt.Println("Postgres version: ", postgresVersion)
	
	var person Person
	if err = db.QueryRow("SELECT id, name, address FROM Person WHERE id = $1", 1).Scan(&person.ID, &person.Name, &person.Address); err != nil {
		panic(err)
	}
	
	fmt.Printf("%+v", person)
}
