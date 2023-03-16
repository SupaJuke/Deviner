package main

import (
	// Local imports
	"fmt"
	"log"

	db "github.com/SupaJuke/pooe-guessing-game/go/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	db.InitDB()

	// db.Db.

	fmt.Println("weh?")
}
