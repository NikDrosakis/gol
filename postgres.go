package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "n130177!"
	dbname   = "gol"
)

func pgcon() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	// close database
	defer db.Close()
	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Pgsql Connected!")
}
