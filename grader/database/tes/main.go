package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "akhsan"
	password = "password"
	dbname   = "dbtest"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// melakukan ping ke database PostgreSQL
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
