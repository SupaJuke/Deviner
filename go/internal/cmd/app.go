package cmd

import (
	"log"

	db "github.com/SupaJuke/pooe-guessing-game/go/internal/database"
	"github.com/SupaJuke/pooe-guessing-game/go/internal/service"
	"github.com/joho/godotenv"
)

// const defaultPort = "8080"

func Run() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	db.InitDB()
	service.Serve()
}
