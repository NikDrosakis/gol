package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Pass      string `json:"pass"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func mariacon() {
	db, err := sql.Open("mysql", "root:n130177!@tcp(localhost:3306)/gs_nikosdrosakisgr?parseTime=true")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT id,name,pass,firstname,lastname FROM user")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.Name, &user.Pass, &user.Firstname, &user.Lastname)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(user.Name + " " + user.Pass + " " + user.Firstname + " " + user.Lastname)
	}

}
