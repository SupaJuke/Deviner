package pooe_hw

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "password"
	DBNAME   = ""
)

func main() string {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
	)
	DB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	return "Why hello there"
}
