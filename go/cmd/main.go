package main

import (
	// Local imports

	"log"

	db "github.com/SupaJuke/pooe-guessing-game/go/internal/database"

	"github.com/joho/godotenv"
)

// const defaultPort = "8080"

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	db.InitDB()

	// router := http.NewServeMux()
	// handler := router.ServeHTTP()
	// router.Handle("/")
}
