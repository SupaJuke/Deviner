package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	
    var (
        USER        = os.Getenv("DB_USER")
        PASSWORD    = os.Getenv("DB_PASS")
        HOST        = os.Getenv("DB_HOST")
        PORT        = os.Getenv("DB_PORT")
        NAME        = os.Getenv("DB_NAME")
    )

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, NAME,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln("Failed connecting to db", err)
    }

	Db = db
}
