package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DbConnect(HOST string, PORT int, USER, POSTGRES_PASSWORD, POSTGRES_DBNAME string) *sql.DB {
	fmt.Println("HOST: ", HOST)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, POSTGRES_PASSWORD, POSTGRES_DBNAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	return db
}
