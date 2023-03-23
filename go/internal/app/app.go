package app

import (
	"log"

	db "github.com/SupaJuke/Deviner/go/internal/database"
	"github.com/SupaJuke/Deviner/go/internal/routes"
	"github.com/joho/godotenv"
)

// const defaultPort = "8080"

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db.InitDB()
	log.Println("Initialized connection to database")
	routes.Serve()
}
