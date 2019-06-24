package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Player struct {
	ID string
	Name string
	Scores []int64
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
	
	var player Player
	if err = db.QueryRow("SELECT id, name, scores FROM ArrayExample WHERE id = $1", 1).Scan(&player.ID, &player.Name, pq.Array(&player.Scores)); err != nil {
		panic(err)
	}
	
	fmt.Printf("%+v", player)
}
