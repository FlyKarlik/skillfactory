package main

import (
	"database_skillfactory/pkg/storage"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := storage.NewPostgresDb(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	store := storage.NewStore(db)
	//Дальнейшая реализация кода

}
