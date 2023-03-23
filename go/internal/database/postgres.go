package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	var (
		USER     = os.Getenv("DB_USER")
		PASSWORD = os.Getenv("DB_PASS")
		HOST     = os.Getenv("DB_HOST")
		PORT     = os.Getenv("DB_PORT")
		NAME     = os.Getenv("DB_NAME")
		SCHEMA   = os.Getenv("DB_SCHEMA")
	)

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, NAME,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln("Failed connecting to db", err)
	}

	DB = db
	stmt := fmt.Sprintf("SET search_path TO %s;", SCHEMA)
	_, err = DB.Exec(stmt)
	if err != nil {
		log.Fatalln("Failed while setting search_path: ", err)
	}
}
