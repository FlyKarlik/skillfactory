package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type user struct {
	id   int
	name string
}

func main() {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=nikitasavin dbname=testing password=Nikita18726 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	//_, err = db.Exec(`INSERT INTO users(name) VALUES ($1)`, "Nikita")
	//if err != nil {
	//	log.Fatal(err)
	//}

	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var u user

		rows.Scan(
			&u.id,
			&u.name,
		)

		fmt.Println(u)
	}

	

}
